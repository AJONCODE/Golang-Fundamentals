Panic is called during a run time error and fatally kill the program

Recover tells Go what to do when that happens
    Returns what was passed to panic.

Recover must be paired with defer, which will fire even after a panic