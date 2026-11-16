

import customtkinter as ctk

class Button(ctk.CTkButton):

    def __init__(self,parent,text,command,color,h_color,t_color="white",style="default",**kwargs):
        super().__init__(
                master=parent,
                text=text,
                fg_color=color,
                command=command,
                hover_color=h_color,
                text_color=t_color,
                font=(ctk.CTkFont(family="Roboto", size=22, weight="bold")),
                compound="left",
                corner_radius=8,
                border_width=0,
                width=200,
                height=38,
                cursor="hand2",
                **kwargs
                )

        
