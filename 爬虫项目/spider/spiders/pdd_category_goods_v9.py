# -*- coding: utf-8 -*-
# 获取首页分类下的产品列表
import scrapy
import json, time, sys, random, pddSign, urllib,setting, pyssdb,string,os,uuid, redis, logging
from spider.items import CategoryGoodsItem
from scrapy.utils.project import get_project_settings

from scrapy.spidermiddlewares.httperror import HttpError
from twisted.internet.error import TimeoutError, TCPTimedOutError, ConnectionRefusedError, DNSLookupError
from twisted.web._newclient import ResponseFailed, ResponseNeverReceived
from scrapy.utils.response import response_status_message # 获取错误代码信息

class PddCategoryGoodsV9Spider(scrapy.Spider):
	name  = 'pdd_category_goods_v9'
	queue_name  = 'queue_pdd_category_goods_list'
	pagesise = 100
	realPagesise = 20

	gateway_url = 'http://gateway.91cyt.com/api/curl'

	custom_settings = {
		# 'LOG_FILE': '',
		# 'LOG_LEVEL': 'DEBUG',
		# 'LOG_ENABLED': True,
		'RETRY_ENABLED': False,
		'DOWNLOAD_TIMEOUT': 5,
		# 'RETRY_TIMES': 20,
		'RETRY_HTTP_CODECS':[403,422,429],
		'DOWNLOAD_DELAY': 0.01,
		'CONCURRENT_REQUESTS': 10
	}

	def __init__(self, hash_num=0, process_nums=1):
		self.ssdb_client = pyssdb.Client(get_project_settings().get('SSDB_HOST'), 8888)
		self.hash_num = int(hash_num)  ##当前脚本号
		self.process_nums = int(process_nums)  ##脚本总数

	def start_requests(self):
		is_end	=	False
		while not is_end:
			category = self.ssdb_client.qpop_front(self.queue_name, 1)

			if not category: ##没有数据返回
				is_end = True
				continue

			detail = json.loads( category.decode('utf-8') )
			if type(detail) != int: ##当值是subject_id 时
				# print (detail);
				cat_id      = detail['subject_id']
				opt_type    = detail['opt_type']
				offset 		= 0 if 'offset' not in detail.keys() else detail['offset']

				uri = self.build_uri(opt_type, cat_id, offset, '')
				meta = {'cat_id':cat_id,'opt_type':opt_type,'offset':offset, 'category':detail}
				# print(url)
				headers = self.make_headers()

				form_data = {
					'name': 'getOperationGroups',
					'method': 'GET',
					'domain': 'http://apiv4.yangkeduo.com',
					'uri': uri,
					'headers': json.dumps(headers),
				}

				logging.debug(json.dumps(form_data))

				yield scrapy.FormRequest(url=self.gateway_url, formdata=form_data, meta=meta, headers={},
									 callback=self.parse, dont_filter=True, errback=self.errback_httpbin)
				# yield scrapy.Request(url, meta=meta, callback=self.parse, headers=headers, dont_filter=True, errback=self.errback_httpbin)

	def parse(self, response):
		# print(response.meta)
		meta = response.meta
		offset = response.meta['offset']
		cat_id = response.meta['cat_id']
		opt_type = response.meta['opt_type']
		receive_info = response.body.decode('utf-8')
		# self.save_goods_log(cat_id, offset, receive_info)
		data = json.loads(receive_info)
		logging.debug(json.dumps({
			'cat_id': cat_id,
			'offset': offset,
			'receive_info': data
		}))
		goods_lists = []
		# print('parse_before', flip, offset,len(data['goods_list']),cat_id)
		if 'goods_list' not in data.keys():
			self.err_after(meta, True)
			return False
		# flip = data['flip']
		if len(data['goods_list']) > 0:
			i = 0
			for goods_data in data['goods_list']:
				i += 1
				rank = offset + i
				goods_data['rank'] = rank
				goods_data['subject_id'] = cat_id
				goods_data['type'] = 2
				goods_lists.append(goods_data)

			offset += i
			item = CategoryGoodsItem()
			item['goods_lists'] = goods_lists
			# print('parse_middle',flip,offset,len(data['goods_list']),cat_id)
			yield item
			# print('parse_after',flip,offset,len(data['goods_list']),cat_id)
			if i >= self.realPagesise and offset < 1000 - self.realPagesise:
				uri = self.build_uri(opt_type, cat_id, offset, '')
				meta['offset'] = offset
				headers = self.make_headers()
				form_data = {
					'name': 'getOperationGroups',
					'method': 'GET',
					'domain': 'http://apiv4.yangkeduo.com',
					'uri': uri,
					'headers': json.dumps(headers),
				}

				logging.debug(json.dumps(form_data))

				yield scrapy.FormRequest(url=self.gateway_url, formdata=form_data, meta=meta, headers={},
									 callback=self.parse, dont_filter=True, errback=self.errback_httpbin)
				# yield scrapy.Request(url, meta=meta, callback=self.parse, headers=headers, dont_filter=True, errback=self.errback_httpbin)

	'''生成headers头信息'''
	def make_headers(self):
		chrome_version   = str(random.randint(59,63))+'.0.'+str(random.randint(1000,3200))+'.94'
		headers = {
			# "Host":"mobile.yangkeduo.com",
			# "Accept":"text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8",
			# "Accept-Language":"zh-CN,zh;q=0.9,en;q=0.8",
			# "Accept-Encoding":"gzip, deflate",
			# "Referer":"http://yangkeduo.com/goods.html?goods_id=442573047&from_subject_id=935&is_spike=0&refer_page_name=subject&refer_page_id=subject_1515726808272_1M143fWqjQ&refer_page_sn=10026",
			# "Connection":"keep-alive",
			# 'User-Agent':'phh_android_version/Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/'+chrome_version+' Safari/537.36',
			#"Host":"mobile.yangkeduo.com",
			#"Accept":"text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8",
			#"Accept-Language":"zh-CN,zh;q=0.9,en;q=0.8",
			#"Accept-Encoding":"gzip, deflate",
			#"Host":"yangkeduo.com",
			"Referer":"",
			"Cookie":'',
			"AccessToken":"",
			#"Connection":"keep-alive",
			'User-Agent':'',
		}

		# ip = str(random.randint(100, 200))+'.'+str(random.randint(1, 255))+'.'+str(random.randint(1, 255))+'.'+str(random.randint(1, 255))
		# headers['CLIENT-IP'] 	=	ip
		# headers['X-FORWARDED-FOR']=	ip
		return headers

	def build_uri(self, opt_type, cat_id, offset, flip):
		pdd_sign = pddSign.pddSign()
		href = 'http://mobile.yangkeduo.com/catgoods.html?opt_id='+str(cat_id)+'&opt_type='+str(opt_type)+'&opt_name=&opt_g=1&refer_page_name=search&refer_page_id=10031_1539251323283_teZmqPRkL8&refer_page_sn=10031&page_id=10028_1539251490140_R7nIY9glyC&is_back=1&sort_type=DEFAULT&act_status=0&opt_index='
		anti_content = pdd_sign.messagePackV2('0al', href)
		url = '/operation/'+str(cat_id)+'/groups?opt_type='+str(opt_type)+'&offset='+str(offset)+'&size='+str(self.pagesise)+'&sort_type=DEFAULT&&pdduid=0&anti_content='+urllib.parse.quote(anti_content)
		# if flip.strip()!='':
		# 	url+='&flip='+flip
		return url

	def get_list_id(self, opt_id):
		uuid_string = str(uuid.uuid1())
		uuid_string = ''.join(uuid_string.split('-'))
		uuid_string = uuid_string.replace("-", "")
		uuid_string = uuid_string[0:10]
		return str(opt_id) + '_' + uuid_string
		# return str(opt_id) + '_rec_list_catgoods_' + ''.join(random.sample(string.ascii_letters + string.digits, 6))

	def errback_httpbin(self, failure):
		request = failure.request
		if failure.check(HttpError):
			response = failure.value.response
			# print( 'errback <%s> %s , response status:%s' % (request.url, failure.value, response_status_message(response.status)) )
			self.err_after(request.meta)
		elif failure.check(ResponseFailed):
			# print('errback <%s> ResponseFailed' % request.url)
			self.err_after(request.meta, True)
 
		elif failure.check(ConnectionRefusedError):
			# print('errback <%s> ConnectionRefusedError' % request.url)
			self.err_after(request.meta, True)
 
		elif failure.check(ResponseNeverReceived):
			# print('errback <%s> ResponseNeverReceived' % request.url)
			self.err_after(request.meta)
 
		elif failure.check(TCPTimedOutError, TimeoutError):
			# print('errback <%s> TimeoutError' % request.url)
			self.err_after(request.meta, True)
		else:
			# print('errback <%s> OtherError' % request.url)
			self.err_after(request.meta)

	def err_after(self, meta, remove = False):
		category = meta['category']
		category['offset'] = meta['offset']
		self.ssdb_client.qpush_back(self.queue_name, json.dumps(category))  # 失败重新放入队列

	def save_goods_log(self, cat_id, page, data):
		date = time.strftime('%Y-%m-%d')

		file_path = '/data/spider/logs/category_log/'+time.strftime('%Y-%m-%d')
		if not os.path.exists(file_path):
			os.makedirs(file_path)

		file_path += '/'+str(cat_id)
		if not os.path.exists(file_path):
			os.makedirs(file_path)

		file_name = file_path+'/'+str(page)+'.log'
		with open(file_name, "a+") as f:
			f.write(data+"\r\n")