"""
标签的三种放置方式
"""
import tkinter as tk


win = tk.Tk()
win.title("my window")
win.geometry("400x600+200+200")

# tk.Label(win, text=1).pack(side="top")
# tk.Label(win, text=1).pack(side="bottom")
# tk.Label(win, text=1).pack(side="left")
# tk.Label(win, text=1).pack(side="right")

# for i in range(4):
#     for j in range(3):
#         tk.Label(win, text=1).grid(row=i, column=j, ipadx=10, ipady=20)

tk.Label(win, text=1).place(x=50, y=100, anchor="nw")

if __name__ == '__main__':
    win.mainloop()
