# git



## stash

暂存本次未修改的内容到栈中

```bash
# 将未完成的修改保存的一个栈上
git stash

# 应用这些改动到本地
git stash apply

# 查看存储
git stash list

# 如果不指定一个储藏，git认为指定的是最近的储藏
git stash apply stash@{2}

# 来应用储藏然后立即从栈上扔掉它
git stash pop

# 家属将要移除的储藏的名字来移除她
git stash drop name
```



## 文件状态

| 状态      | 中文名               | 说明 |
| --------- | -------------------- | ---- |
| Untracked | 未跟踪               |      |
| Unmodify  | 文件已经入库，未修改 |      |
| Modified  | 文件已修改           |      |
| Staged    | 暂存状态             |      |

