import tkinter as tk
import pickle

win = tk.Tk()
win.title("welcome to Boluce Python")
win.geometry("500x350+200+200")

# welcome image
# canvas = tk.Canvas(win, height=200, width=500)
# image_file = tk.PhotoImage(file="welcome.gif")
# image = canvas.create_image(0, 0, anchor="nw", image=image_file)
# canvas.pack(side="top")
tk.Label(win, text="User Name:").place(x=50, y=150)
tk.Label(win, text="Password").place(x=50, y=190)

var_usr_name = tk.StringVar()
var_usr_name.set("examplae@python.com")

entry_user_name = tk.Entry(win, textvariable=var_usr_name)
entry_user_name.place(x=160, y=150)

var_usr_pwd = tk.StringVar()
entry_user_name = tk.Entry(win, textvariable=var_usr_pwd, show="*")
entry_user_name.place(x=160, y=190)


def usr_login():
    usr_name = var_usr_name.get()
    usr_pwd = var_usr_pwd.get()


def usr_sign_up():

    def sign_to_Mofan_Python():
        np = new_pwd.get()
        npf = new_pwd_confirm.get()
        nn = new_name.get()

    window_sign_up = tk.Toplevel(win)
    window_sign_up.geometry("350x200")
    window_sign_up.title("Sign up window")

    new_name = tk.StringVar()
    new_name.set("example@python.com")
    tk.Label(window_sign_up, text="User name:").place()


btn_login = tk.Button(win, text="Login", command=usr_login)
btn_login.place(x=170, y=230)

btn_sign_up = tk.Button(win, text="Sign up", command=usr_sign_up)
btn_sign_up.place(x=250, y=230)


if __name__ == '__main__':
    win.mainloop()