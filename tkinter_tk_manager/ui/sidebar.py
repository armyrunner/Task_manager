
"""
The sidebar contains the dropdown menu for tasks,
and the add,update and delete buttons for the actions for 
tasks.

"""

import customtkinter as ctk
from ui.components.button import Button
from ui.components.dropdown import DropDown
from pytablericons import TablerIcons, OutlineIcon
from ui.components.add_model import AddModel
from ui.components.update_model import UpdateModel

BUTTON_STYLES = {
    "primary": {
        "color": "#3498DB",
        "h_color": "#2980B9"
    },
    "success": {
        "color": "#27AE60",
        "h_color": "#1E8449"
    },
    "danger": {
        "color": "#E74C3C",
        "h_color": "#C0392B"
    },
    "exit":{
        "color": "#E74C3C",
        "h_color": "#C0392B",
    },
    "settings":{
        "settings": "#27AE60",
        "h_color": "#1E8449",
    },
    "home":{
        "home": "#3498DB",
        "h_color": "#2980B9",
    },
}

class SideBar:

# intialize sidebar

    def __init__(self,parent,content_frame):

        self.parent = parent
        self.content_frame=content_frame

        # create sibebar frame
        self.frame = ctk.CTkFrame(
                self.parent,
                width=220,
                corner_radius=1,
                fg_color="#1E1E2E"
                )

        self.frame.grid(
                row=0,
                column=0,
                sticky="nsew"
                )
        self.frame.grid_propagate(False)

        #Horizontal bar at the top of the sidebar
        self.logo_frame = ctk.CTkFrame(
                self.frame,
                height=80,
                corner_radius=0,
                fg_color="#364456"
                )
        
        self.logo_frame.grid(
                row=0,
                column=0,
                sticky="ew",
                padx=10,
                pady=(10,0)
                )

        self.logo_frame.grid_propagate(False)

        # Sidebar grid configuration
        self.frame.grid_columnconfigure(1, weight=1)

        # Allow row 5 to expand and push Exit to the bottom
        self.frame.grid_rowconfigure(5, weight=1)
        
        self.plusIcon = ctk.CTkImage(
            light_image=TablerIcons.load(
                OutlineIcon.PLUS, size=24, color="#FFFFFF"
            ),
            dark_image=TablerIcons.load(
                OutlineIcon.PLUS, size=24, color="#FFFFFF"
            ),
            size=(24, 24)
        )

        self.pencilIcon = ctk.CTkImage(
            light_image=TablerIcons.load(
                OutlineIcon.PENCIL, size=24, color="#FFFFFF"
            ),
            dark_image=TablerIcons.load(
                OutlineIcon.PENCIL, size=24, color="#FFFFFF"
            ),
            size=(24, 24)
        )

        self.deleteIcon = ctk.CTkImage(
            light_image=TablerIcons.load(
                OutlineIcon.TRASH, size=24, color="#FFFFFF"
            ),
            dark_image=TablerIcons.load(
                OutlineIcon.TRASH, size=24, color="#FFFFFF"
            ),
            size=(24, 24)
        )

        self.exitIcon = ctk.CTkImage(
            light_image=TablerIcons.load(
                OutlineIcon.LOGOUT, size=24, color="#FFFFFF"
            ),
            dark_image=TablerIcons.load(
                OutlineIcon.LOGOUT, size=24, color="#FFFFFF"
            ),
            size=(24, 24)
        )

        self.add_button = Button(
                parent=self.frame,
                text="Add",
                command=self.on_add,
                image=self.plusIcon,
                **BUTTON_STYLES["primary"]
                )

        self.update_button = Button(
                parent=self.frame,
                text="Update",
                command=self.on_update,
                image=self.pencilIcon,
                **BUTTON_STYLES["success"]
                )

        self.delete_button = Button(
                parent=self.frame,
                text="Delete",
                command=self.on_delete,
                image=self.deleteIcon,
                **BUTTON_STYLES["danger"]
                )
        
        self.exit_button = Button(
                parent=self.frame,
                text="EXIT",
                command=self.exitManager,
                image=self.exitIcon,
                **BUTTON_STYLES["exit"]
                )
        
        self.current_task = DropDown(
            parent=self.frame,
            values=["Honey Due", "Washing Machine", "Homework"],
            command=self.listTask,
            fg_color="#FFFFFF",
            text_color="#000000",
            button_color="#D3D3D3",
            button_hover_color="#BEBEBE",
            dropdown_fg_color="#FFFFFF",
            dropdown_text_color="#000000"
        )

        self.current_task.grid(row=1,column=0,padx=5,pady=15)
        self.add_button.grid(row=2,column=0,padx=5,pady=10)
        self.update_button.grid(row=3,column=0,padx=5,pady=10)
        self.delete_button.grid(row=4,column=0,padx=5,pady=10)

        self.exit_button.grid(row=6,column=0,padx=5,pady=10,sticky="ew")

# create category dropdown (load existing, select current, add new categories

# create all tasks navitagtion
# create completed naviagation

    # create add task button
    def on_add(self):
        self.add_model = AddModel(self.content_frame)

        if hasattr(self, "add model"):
            self.add_model.destroy()

        self.add_model.grid(row=0, column=0, sticky="nsew", padx=20,pady=20)


    # create update task button

    def on_update(self):
        self.update_model = UpdateModel(self.content_frame)

        if hasattr(self, "update model"):
            self.update_model.destroy()

        self.update_model.grid(row=0, column=0, sticky="nsew", padx=20,pady=20)


    # create delete task button

    def on_delete(self):
        print("delete a task")

    def listTask(self):
        print("Selected Task")

    def exitManager(self):
        self.parent.destroy()

# place settings at bottom
# place logout button under settings
# end initialize

# When categories change: (Notify main window of selected category)
# When all tasks is clicked: (Display existing task list)
# When completed task is clicked: (display completed task list)
# When add is clicked: (add existing add_task function)
# when update is clicked: (call existing update function - and open update view)
# when delete is clicket: (call existing detete function - and open delete view)
# when settings is clicked: ( open settings view) ** future implementation **
# when logout is clicked: ( run logout action) ** Future implementation**

# End of Class
 

