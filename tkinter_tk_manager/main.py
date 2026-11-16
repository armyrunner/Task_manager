

from ui.app import App


SCREEN_WIDTH = 1280
SCREEN_HEIGHT = 768

def main():
    app = App(SCREEN_HEIGHT,SCREEN_WIDTH,"Task Manager",200,100)
    app.run()

if __name__ == "__main__":
    main()
