

import customtkinter as ctk
from ui.components.button import Button
from ui.components.dropdown import DropDown
from pytablericons import TablerIcons, OutlineIcon
from ctkdateentry import CTkDateEntry,CTkStringVar

BUTTON_STYLES = {
    "save": {
        "color": "#3498DB",
        "h_color": "#2980B9"
    },
}

class AddModel(ctk.CTkFrame):



    def __init__(self,parent,**kwargs):
        super().__init__(
            master=parent,
            fg_color=["#F2F2F2", "#2B2B2B"],  # Card background
            bg_color= ["#F2F2F2", "#2B2B2B"],
            border_color=["#D3D3D3", "#404040"],
            border_width=2,
            corner_radius=12,
            **kwargs
        )
        self.grid_columnconfigure(0, weight=1, uniform="columns")
        self.grid_columnconfigure(1, weight=1, uniform="columns")

        # Main Title and location of Model
        self.frame_title = ctk.CTkLabel(
            self,
            text="ADD TASK", 
            font=(ctk.CTkFont(family="Segoe UI", size=18, weight="bold"))
        )
        
        self.frame_title.grid(row=0,column=0, columnspan=6, padx=20, pady=(20,15) )
        
        """

        ******** LEFT FRAME SIDE **********

        """
        self.left_frame = ctk.CTkFrame(
            self,
            width=500,
            fg_color = "transparent"
        )
        self.left_frame.grid(row=1, column=0,padx=(20, 10), pady=10, sticky="nsew")
        self.left_frame.grid_columnconfigure(0,weight=1)

        # Add Task Entry
        self.task_entry_label = ctk.CTkLabel(
            self.left_frame,
            text="Task Description", 
            font=(ctk.CTkFont(size=12, weight="bold"))
        )
        self.task_entry_label.grid(row=0,column=0,padx=5, pady=5, sticky="w" )

        self.task_entry =  ctk.CTkEntry(
            self.left_frame,
            placeholder_text="Enter Task Description...",
        )                                  
        self.task_entry.grid(row=1,column=0, padx=5, pady=(0,15), sticky="ew" )

        # Note: Text entry to describe what the progress 
        self.note_entry_label = ctk.CTkLabel(
            self.left_frame,
            text="Notes", 
            font=(ctk.CTkFont(size=12, weight="bold"))
        )
        self.note_entry_label.grid(row=2,column=0, padx=5, pady=5, sticky="w" )

        self.note_entry =  ctk.CTkEntry(
            self.left_frame, 
            placeholder_text="Notes Section...",
        )                                  
        self.note_entry.grid(row=3,column=0,padx=5, pady=(0,15), sticky="ew" )

        """

        ******** RIGHT FRAME SIDE **********

        """
        self.right_frame = ctk.CTkFrame(
            self,
            width=300,
            fg_color = "transparent",

        )
        self.right_frame.grid(row=1, column=1, padx=(10, 20), pady=10, sticky="nsew")
        self.right_frame.grid_columnconfigure(0,weight=1)


        # Due Date and Start Date Information Label and date picker
        self.due_date_var = CTkStringVar(self.right_frame, value="Enter a Date")
        self.due_date_label = ctk.CTkLabel(
            self.right_frame, 
            text="Due Date", 
            font=ctk.CTkFont(size=12,weight="bold")
        )

        self.due_date_entry = CTkDateEntry(
            self.right_frame,
            variable=self.due_date_var,
            justify="left",
            font=ctk.CTkFont(family="Segoe UI",size=12),
            width=250
        )
        
        self.due_date_label.grid(row=0, column=0, padx=5, pady=(0, 5), sticky="w")
        self.due_date_entry.grid(row=1, column=0, padx=5, pady=(0, 10), sticky="ew")
        
        self.start_date_var = CTkStringVar(self.right_frame, value="Enter a Date")
        self.start_date_label = ctk.CTkLabel(
            self.right_frame, 
            text="Start Date", 
            font=ctk.CTkFont(size=12,weight="bold")
        )

        self.start_date_entry = CTkDateEntry(
            self.right_frame, 
            variable=self.start_date_var,
            justify="left",
            font=ctk.CTkFont(family="Segoe UI",size=12),
            width=250
        )

        self.start_date_label.grid(row=2, column=0, padx=5, pady=(0, 5), sticky="w")
        self.start_date_entry.grid(row=3, column=0, padx=5, pady=(0, 10), sticky="ew")

       # Finished Date, Category, status Section Information
        self.finish_date_label = ctk.CTkLabel(
            self.right_frame,  
            text="Finish Date", 
            font=ctk.CTkFont(size=12,weight="bold")
        )

        self.category_label = ctk.CTkLabel(
            self.right_frame,  
            text="Category", 
            font=ctk.CTkFont(size=12,weight="bold")
        )

        self.finish_date_var = CTkStringVar(self.right_frame, value="Enter a Date")

        self.finish_date_entry = CTkDateEntry(
            self.right_frame, 
            variable=self.finish_date_var,
            justify="left",
            font=ctk.CTkFont(family="Segoe UI",size=12),
            width=250
            )
        self.finish_date_label.grid(row=4, column=0, padx=5, pady=(0, 5), sticky="w")
        self.finish_date_entry.grid(row=5, column=0, padx=5, pady=(0, 10), sticky="ew")

        self.category_var = CTkStringVar(self.right_frame, value="")
        self.category_menu = DropDown(
            self.right_frame, 
            variable=self.category_var,
            values=["Honey Due", "Washing Machine", "Homework"],
            command=self.listMenu,
            fg_color="#FFFFFF",
            text_color="#000000",
            button_color="#D3D3D3",
            button_hover_color="#BEBEBE",
            dropdown_fg_color="#FFFFFF",
            dropdown_text_color="#000000"
        )

        self.category_label.grid(row=6, column=0, padx=5, pady=(0, 5), sticky="w")
        self.category_menu.grid(row=7, column=0, padx=5, pady=(0, 10), sticky="ew")

        self.status_label = ctk.CTkLabel(
            self.right_frame, 
            text="Status",
            font=ctk.CTkFont(size=12, weight="bold")
        )
  

        self.status_var = ctk.StringVar(
            self.right_frame,
            value="Not Started"
        )

        self.status_menu = ctk.CTkOptionMenu(
            self.right_frame,
            variable=self.status_var,
            values=[
                "Not Started",
                "In Progress",
                "On Hold",
                "Completed",
                "Cancelled"
            ],
            height=35
        )
        self.status_label.grid(row=8, column=0, padx=5, pady=(0, 5), sticky="w")
        self.status_menu.grid(row=9, column=0, padx=5, pady=(0, 10), sticky="ew")


        # Save button, action and icon

        self.saveIcon = ctk.CTkImage(
            light_image=TablerIcons.load(
                OutlineIcon.DEVICE_FLOPPY, size=24, color="#FFFFFF"
            ),
            dark_image=TablerIcons.load(
                OutlineIcon.DEVICE_FLOPPY, size=24, color="#FFFFFF"
            ),
            size=(24, 24)
        )

        self.save_button = Button(
                self.right_frame,
                text="SAVE",
                command=self.save_task,
                image=self.saveIcon,
                **BUTTON_STYLES["save"]
                )
        
        self.save_button.grid(row=10,column=0, columnspan=2,padx=20,pady=(5,20),sticky="e")

    def listMenu(self, choice=None):
        print("Category Menu Items:")
        
    def save_task(self):
        task_name = self.task_entry.get()
        due_date = self.due_date_var.get()
        start_date = self.start_date_var.get()
        finish_date = self.finish_date_var.get()
        category = self.category_var.get()
        status = self.status_var.get()
        notes = self.note_entry.get()

        print(f"Task: {task_name}")
        print(f"Due Date: {due_date}")
        print(f"Start Date: {start_date}")
        print(f"Finish Date: {finish_date}")
        print(f"Category: {category}")
        print(f"Status: {status}")
        print(f"Notes: {notes}")

