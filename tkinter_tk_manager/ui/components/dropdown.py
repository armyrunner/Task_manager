


import customtkinter as ctk


class DropDown(ctk.CTkOptionMenu):

    def __init__(self, parent, values, command, **kwargs):
        super().__init__(
            master=parent,
            values=values,
            command=command,
            width=190,
            height=38,
            corner_radius=8,
            **kwargs
        )


