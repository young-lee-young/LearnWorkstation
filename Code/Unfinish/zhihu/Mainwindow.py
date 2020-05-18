import tkinter
import Login
from PIL import ImageTk, Image


class Application(object):
    def __init__(self, master=None):
        self.username = tkinter.StringVar()
        self.password = tkinter.StringVar()
        self.root = master
        self.init_window()
        self.init_frame()
        self.root.mainloop()

    def init_window(self):
        self.root.title('知乎推荐系统')
        self.root.geometry('800x500+100+100')

        # self.root['bg'] = background_pic
        self.root.resizable(False, False) # 禁止调整窗口大小

    def init_frame(self):
        background_pic = tkinter.PhotoImage(file='pic/zhihu.png')
        frame = tkinter.Frame(self.root, text="我是内容,\n请你阅读", justify=tkinter.LEFT, image=background_pic, compound=tkinter.CENTER,
                    font=("华文行楷", 20), fg="white")
        # image = Image.open(r'E:\zhihu.gif')


        # tkinter.Label(frame, text='学号:', image=background_pic)
        # tkinter.Entry(self.page, textvariable=self.username).grid(row=0, column=1, rowspan=1, columnspan=2)
        #
        # tkinter.Label(self.page, text='密码:', pady=5).grid(row=1, column=0)
        # tkinter.Entry(self.page, textvariable=self.password).grid(row=1, column=1, rowspan=1, columnspan=2)
        #
        # tkinter.Button(self.page, text='登录', pady=5, command=self.logincheck).grid(row=3, column=0)
        # tkinter.Button(self.page, text='退出', pady=5, command=self.logout).grid(row=3, column=2)
        frame.pack()

        # def logincheck(self):
        #     username = self.username.get()
        #     password = self.password.get()
        #     session, response_code = Login.main(username, password)
        #     if response_code != 200:
        #         print(response_code)
        #     else:
        #         self.page.destroy()
        #         Info.Info(session, self.root)
        #
        # def logout(self):
        #     self.root.destroy()