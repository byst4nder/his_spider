"""
Frame 框架
"""
import tkinter as tk


win = tk.Tk()
win.title("my window")
win.geometry("400x600+200+200")
tk.Label(win, text="on the window").pack()

frm = tk.Frame(win)
frm.pack()
frm_1 = tk.Frame(frm,)
frm_r = tk.Frame(frm)
frm_1.pack(side="left")
frm_r.pack(side="right")

tk.Label(frm_1, text="on the frm_l1").pack()
tk.Label(frm_1, text="on the frm_l2").pack()
tk.Label(frm_1, text="on the frm_l3").pack()
tk.Label(frm_r, text="on the frm_r1").pack()

if __name__ == '__main__':
    win.mainloop()