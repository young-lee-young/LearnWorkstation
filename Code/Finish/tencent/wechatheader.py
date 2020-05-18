import itchat
import os
import math
from PIL import Image


def login():
    itchat.auto_login(enableCmdQR=False)


def get_header():
    num = 0
    friends = itchat.get_friends(update=True)
    for friend in friends:
        img = itchat.get_head_img(userName=friend['UserName'])
        with open('D:/spider/wechat/' + str(num) + '.png', 'wb') as f:
            f.write(img)
        num += 1


def image():
    pictures = os.listdir('D:/spider/wechat')
    eachsize = int(math.sqrt(float(640 * 640) / len(pictures)))
    numline = int(640 / eachsize)
    toimage = Image.new('RGBA', (640, 640))
    x = 0
    y = 0
    for picture in pictures:
        try:
            img = Image.open('D:/spider/wechat/' + picture)
            img = img.resize((eachsize, eachsize), Image.ANTIALIAS)
            toimage.paste(img, (x * eachsize, y * eachsize))
            x += 1
            if x == numline + 1:
                x = 0
                y += 1
        except IOError:
            print('IOError')
    toimage.save('D:/spider/wechat/toimage.png')


def main():
    login()
    get_header()
    image()


if __name__ == '__main__':
    main()