
import customtkinter as ctk
from ui.components.button import Button
class AutoCompleteDropDown(ctk.CTkFrame):

    def __init__(self, parent, variable=None, 
                suggestions=None, command=None,placeholder_text="Search Tasks...",max_suggestion=8, **kwargs):
        super().__init__(parent,fg_color="transparent",**kwargs)

        self.suggestions = list(suggestions or [])
        self.command = command
        self.max_suggestions = max_suggestion

        self.filtered_suggestions = []
        self.active_index = -1
        self.selected_value = None

        self.after_id = None
        self.hide_after_id = None

        self.variable = variable

        # ************************* SEARCH ENTRY **************************

        self.entry = ctk.CTkEntry(
            self,
            textvariable= self.variable,
            placeholder_text=placeholder_text,
            height=35
        )

        self.entry.grid(row=0,column=0,sticky="ew")

        self.grid_columnconfigure(0,weight=1)

        # ************************* SUGGESTIONS POPUP **************************

        self.popup = ctk.CTkToplevel(self)
        self.popup.withdraw()
        self.popup.overrideredirect(True)

        self.popup_frame = ctk.CTkScrollableFrame(
            self.popup,
            width=300,
            height=180,
            fg_color=["#FFFFFF", "#252525"],
            corner_radius=8
        )

        self.popup_frame.grid(row=0, column=0, sticky="nsew")

        self.popup.grid_columnconfigure(0,weight=1)
        self.popup.grid_rowconfigure(0,weight=1)
        self.popup_frame.grid_columnconfigure(0,weight=1)

        # ************************* EVENTS BINDINGS **************************

        self.variable.trace_add(
            "write",
            self._on_text_changed
        )

        self.entry.bind(
            "<Down>",
            self._move_down
        )

        self.entry.bind(
            "<Up>",
            self._move_up
        )

        self.entry.bind(
            "<Return>",
            self._confirm_selection
        )

        self.entry.bind(
            "<Escape>",
            self._hide_popup
        )

        self.entry.bind(
            "<FocusOut>",
            self._schedule_hide
        )

        self.bind(
            "<Destroy>",
            self._on_destroy
        )
    
    # ==================================================
    # PUBLIC METHODS
    # ==================================================

    # Return the text currently in teh search entry
    def get(self):
        return self.variable.get()
    
    # Return the selected task, or None
    def get_selected(self):
        return self.selected_value.get()
    
    # Set the displayed task text
    def set(self,value):
        self.selected_value = None
        self.selected_value = value
        self.variable.set("")
        self._hide_popup()

    # Clear the search entry and selected task
    def clear(self):
        self.selected_value = None
        self.variable.set("")
        self._hide_popup()

    # Replace the availabe task suggestons
    def set_suggestions(self,suggestions):
        self.suggestions = list(suggestions)
        self._on_text_changed()

    # ==================================================
    # SEARCH AND FILTERING
    # ==================================================


    def _on_text_changed(self, *_):
        self.selected_value = None
        self.active_index = -1

        if self.after_id is not None:
            self.after_cancel(self.after_id)
            self.after_id = None

        query = self.get_variable.get().strip().casefold()

        if not query:
            self._hide_popup()
            return
        
        # Small delay avoids filiter on every keystroke
        self.after_id = self.after(
            100,
            lambda q=query: self._filter(q)
        )

    def _filter(self,query):
        self.after_id = None

        # Ignore outdated search queries
        if query != self.get_variable.get().strip().casefold():
            return
        
        self.filtered_suggestions = [
            task
            for task in self.suggestions
            if query in str(task).casefold()
        ][:self.max_suggestions]

        self._clear_popup_items()
        self.active_index = -1

        if not self.filtered_suggestions:
            no_results = ctk.CTkLabel(
                self.popup_frame,
                text = " No matching tasks",
                anchor = "w"
            )

            no_results.grid(row=0,column=0,padx=8,pady=8, sticky="ew")

            self._show_popup()
            return
        for index, task in enumerate(self.filtered_suggestions):
            suggestion_button = Button(
                self.popup_frame,
                text=str(task),
                anchor="w",
                height=32,
                fg_color="transparent",
                hover_color=["#E6E6E6", "#3A3A3A"],
                text_color=["#222222", "#EEEEEE"],
                command=lambda i=index: self._select(i)
            )

            suggestion_button.grid(row=index,column=0,padx=3,pady=1,sticky="ew")

            self._show_popup()

    def _clear_popup_itesm(self):
        for widget in self.popup_frame.winfo_children():
            widget.destroy()

    # ==================================================
    # POPUP POSITION AND VISIBILITY
    # ==================================================

    def _show_popup(self):
        if not self.winfo_exists():
            return
        
        self.update_idletasks()
        x = self.entry.winfo_rootx()
        y = (self.entry.winfo_rooty()+self.entry.winfo_height())

        width = self.entry.winfo_width()

        item_count = max(1,len(self.filtered_suggestions))

        height = min(220,item_count*36 + 8)

        self.popup.geometry(f"{width}x{height}+{x}+{y}")

        self.popup.deiconify()
        self.popup.lift()

    def _hide_popup(self, _event=None):
        if self.hide_after_id is not None:
            try:
                self.after_cancel(self.hide_after_id)
            except ctk.TclError:
                pass
                
            self.popup.withdraw()

        return "break"
    
    def _schedule_hide(self):
        if self.hide_after_id is not None:
            self.after_cancel(self.hide_after_id)
        
        self.hide_after_id = self.after(
            150, 
            self.hide_if_unclosed
        )

    def _hide_if_unfocused(self):
        self.hide_after_id = None

        try:
            focused = self.focus_get()

            if focused is None:
                self._hide_popup()
                
                return

            if focused.winfo_toplevel() not in (self.winfo_toplevel(),self.popup):
                self._hide_popup()
            
        except ctk.TclError:
            self._hide_popup()

    # ==================================================
    # KEYBOARD NAVIGATION
    # ==================================================
    
    def _move_down(self):
        if not self.filtered_suggestions:
            return "break"
        
        self.active_index = min(
            self.active_index + 1,
            len(self.filtered_suggestions - 1)
        )

        self._highlight_active()

        return "break"
    
    def _move_up(self):
        if not self.filtered_suggestions:
            return "break"
        
        self.active_index = max(
            self.active_index -1, 
            0
        )

        self._highlight_active()

        return "break"
    
    def _highlight_active(self):
        for index, widget in enumerate(self.popup_frame.winfo_children()):
            if not isinstance(widget, ctk.CTkButton ):
                continue

            if index == self.active_index:
                widget.configure(
                    fg_color = ["#D9EAF7","#3A5063"]
                )

            else:
                widget.configure(
                    fg_color="transparent"
                )
        
    def _confirm_selection(self,_event=None):
        if(0 <= self.active_index < len(self.filtered_suggestions)):
            self.select(self.active_index)

        elif len(self.filtered_suggestions) == 1:
            self._select(0)

        return "break"
    
    # ==================================================
    # SELECT TASK
    # ==================================================

    def _select(self,index):
        if not (0 <= index < len(self.filtered_suggestions)):
            return
        
        task = self.filtered_suggestions[index]

        self.selected_value = task
        self.variable.set(str(task))

        self._hide_popup()

        if self.command is not None:
            self.command(task)

        self.entry.focus_set() 

    # ==================================================
    # CLEANUP
    # ==================================================

    def _on_destroy(self, event):
        if event.widget is not self:
            return
        
        for callback_id in (self.after_id,self.hide_after_id):
            if callback_id is None:
                try:
                    self.after_cancel(callback_id)
                except ctk.TclError:
                    pass
        try:
            self.popup.destroy()
        except ctk.TclError:
            pass
    



                
    