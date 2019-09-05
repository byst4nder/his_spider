import smtplib

my_sender = '865466036@qq.com'  # 发件人邮箱账号
my_pass = 'tiwypblkspxqbdfi'  # 发件人邮箱密码
my_user = ["abc@venpoo.com","def@venpoo.com","hij@venpoo.com","kkk@venpoo.com","sss@venpoo.com"] # 收件人邮箱账号

class Config:
    USERNAME = 'root'
    PASSWORD = '*****'
    HOSTNAME = '10.24.151.7'
    DATABASE = '***'


def mail():
    try:
        server = smtplib.SMTP_SSL("smtp.qq.com", 465)  # 发件人邮箱中的SMTP服务器，端口是465
        server.login(my_sender, my_pass)  # 括号中对应的是发件人邮箱账号、邮箱密码
        server.quit()  # 关闭连接
    except Exception as e:  # 如果 try 中的语句没有执行，则会执行下面的 ret=False
        ret = False
        return ret


if __name__ == '__main__':

    ret = mail()
    if ret:
        print("邮件发送成功")
    else:
        print("邮件发送失败")