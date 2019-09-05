import tkinter as tk


win = tk.Tk()
win.title("my window")
win.geometry("400x600+200+200")

l = tk.Label(win, text="", background="yellow")
l.pack()

counter = 0


def do_job():
    global counter
    l.config(text="do" + str(counter))
    counter += 1


menubar = tk.Menu(win)
filemenu = tk.Menu(menubar, tearoff=0)
menubar.add_cascade(label="File", menu=filemenu)
filemenu.add_command(label="New", command=do_job)
filemenu.add_command(label="Open", command=do_job)
filemenu.add_command(label="Save", command=do_job)
filemenu.add_separator()
filemenu.add_command(label="Exit", command=win.quit)

submenu = tk.Menu(filemenu)
filemenu.add_cascade(label="Import", menu=submenu, underline=0)
submenu.add_command(label="Submenul", command=do_job)
submenu.add_command(label="Submenu2", command=do_job)

editmenu = tk.Menu(menubar, tearoff=0)
menubar.add_cascade(label="Edit", menu=editmenu)
editmenu.add_command(label="Cut", command=do_job)
editmenu.add_command(label="Copy", command=do_job)
editmenu.add_command(label="Paste", command=do_job)
win.config(menu=menubar)


if __name__ == '__main__':
    win.mainloop()