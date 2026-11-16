"""
Program File: app.py

Gathers all the other files (sidebar, task_table) with
components dropdown,button for the side bar

In a nut shell this file owns the window, main loop, and 
any event loop.
        
class App ( main window)
    Initializxe
        create main widow
        configure window size - possiblly global to make it more dynamic for resizing

        create left sidebar frame
        create right sidebar frame

        create sidebar inside left frame - brought in by sidebar.py
        create task_view inside right frame - brought in by task_view.py
    end Initialize

end class

"""

import customtkinter as ctk
from ui.sidebar import SideBar

class App:
    
    def __init__(self,height,width,title,xoffset,yoffset):
        self.screen_height = height
        self.screen_width = width
        self.title = title
        self.xoffset = xoffset
        self.yoffset = yoffset
        
    # create and configure the window ( allow for resizing, took sidebar to left, allow content to expand)
        # configure the appearance of the customTkinter 
        # setup the overall layout
        # create sidebar on the leftside
        # create the task view on the rightside

        ctk.set_appearance_mode("light")
        ctk.set_default_color_theme("blue")
        
        # create and configure the main window
        self.root = ctk.CTk()
        self.root.title(self.title)
        self.root.geometry(f"{self.screen_width}X{self.screen_height}"
                      f"{self.xoffset}+{self.yoffset}")
         
        self.root.minsize(1280,780)

        self.root.resizable(False,False) 
        # configuration of the main window
        self.root.grid_rowconfigure(0,weight=1)
        self.root.grid_columnconfigure(1,weight=1)

        # leftside of the main frame - from main action sidebar implmentation is in sidebar.py
        self.sidebar_frame = ctk.CTkFrame(
                self.root,
                corner_radius=0,
                width=200,
                fg_color="#202936"
                )
        self.sidebar_frame.grid(
                row=0,
                column=0,
                sticky="nsew"
                )
        
        self.sidebar_frame.grid_propagate(False)

        # main content frame area 
        self.content_frame = ctk.CTkFrame(
                self.root,
                corner_radius=0,
                fg_color="#F4F6F8"
                )
        
        self.content_frame.grid(
                row=0,
                column=1,
                sticky="nsew"
                )


        self.content_frame.grid_rowconfigure(0,weight=1)
        self.content_frame.grid_columnconfigure(1,weight=1)

        self.sidebar = SideBar(self.root,self.content_frame)

    def run(self):

        self.root.mainloop()

  
