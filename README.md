### Pretty CLI todo app in Go
A beautiful command line application written in Go.


```
╔═══╤═══════════════════════╤═══════╤═══════════════════╤═══════════════════╗
║ # │         Task          │ Done? │    Created at     │   Completed at    ║
╟━━━┼━━━━━━━━━━━━━━━━━━━━━━━┼━━━━━━━┼━━━━━━━━━━━━━━━━━━━┼━━━━━━━━━━━━━━━━━━━╢
║ 1 │ ✅ Go shopping        │ Yes   │ 07 JAN 2022 13:15 │ 07 JAN 2022 15:45 ║
║ 2 │ Read the news         │ No    │ 07 JAN 2022 15:28 │ Not yet completed ║
║ 3 │ Go to the dentist     │ No    │ 07 JAN 2022 16:05 │ Not yet completed ║
║ 4 │ ✅ Write some Go code │ Yes   │ 07 JAN 2022 16:06 │ 07 JAN 2022 16:06 ║
║ 5 │ Watch T.V.            │ No    │ 07 JAN 2022 16:14 │ Not yet completed ║
╟━━━┼━━━━━━━━━━━━━━━━━━━━━━━┼━━━━━━━┼━━━━━━━━━━━━━━━━━━━┼━━━━━━━━━━━━━━━━━━━╢
║                        You have 3 pending task(s)                         ║
╚═══╧═══════════════════════╧═══════╧═══════════════════╧═══════════════════╝
```

To add a task
```
go run ./cmd/todo -add
```

To list all tasks
```
make list
```

To mark a task as complete
```
go run ./cmd/todo -complete={# of the task}
```

To delete a task
```
go run ./cmd/todo -del={# of the task}
```
