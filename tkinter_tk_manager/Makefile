
PYTHON = python3
PIP = $(PYTHON) -m pip

.PHONY: run install test clean

run:
	$(PYTHON) main.py

install:
	$(PIP) install customtkinter

test:
	$(PYTHON) -m pytest

clean:
	find . -type d -name __pycache__ -exec rm -rf {} +

