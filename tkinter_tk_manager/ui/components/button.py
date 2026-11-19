

import customtkinter as ctk

class Button(ctk.CTkButton):

    def __init__(
        self, 
        parent, 
        text,
        command,
        color,
        h_color,
        t_color="white",
        style="default",
        **kwargs
    ):

        kwargs.setdefault(
            "font",
            ctk.CTkFont(family="Roboto", size=22, weight="bold")
        )

        kwargs.setdefault("compound", "left")
        kwargs.setdefault("corner_radius", 8)
        kwargs.setdefault("border_width", 0)
        kwargs.setdefault("width", 200)
        kwargs.setdefault("height", 38)
        kwargs.setdefault("cursor", "hand2")

        super().__init__(
            master=parent,
            text=text,
            fg_color=color,
            command=command,
            hover_color=h_color,
            text_color=t_color,
            **kwargs
        )

        
