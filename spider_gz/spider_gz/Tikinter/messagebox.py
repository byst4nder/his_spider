import tkinter as tk
from tkinter import messagebox

win = tk.Tk()
win.title("my window")
win.geometry("400x600+200+200")


def hit_me():
    # tk.Message.info(title="Hi", message=="hahahaha")
    # messagebox.showinfo(title="Hi", message="hahahaha")
    # messagebox.showwarning(title="Hi", message="hahahha")
    # messagebox.showerror(title="Hi", message="No!! never")
    # print(messagebox.askquestion(title="Hi", message="hahahaha"))  return "yes" , "no"
    messagebox.askyesno(title="Hi", message="hahahaha") # rerurn 'yes', "no"
    messagebox.askokcancel(title="Hi", message="hahaha") # return True , "false"


tk.Button(win, text="hit me", command=hit_me).pack()

if __name__ == '__main__':
    win.mainloop()