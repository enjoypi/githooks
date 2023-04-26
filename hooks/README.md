# git hooks

## 目的

通过git仓库的hook各种来做代码及资源提交的基本检查。

希望能够最大限度的避免git并行开发模式引起的人为错误。

## 使用方法

原理：将需要的hook文件拷贝到.git/hooks/目录里，git执行各种操作的时候就会自动触发hook命令

1. 执行hooks\install-pre-commit.bat
2. 修改hooks\hooks.yaml，具体参见各个具体hook的说明，例如：[pre-commit](#pre-commit)

## pre-commit

提交前的检查，调用verify-commit.exe检查当前分支是否允许提交这些文件类型，通过扩展名来判断文件类型。

* 配置格式参见hooks.yaml里的pre-commit部分
```yaml
# 配置当前提交在特定分支上允许的文件类型，不需要检查的分支不用配
pre-commit:
  # 需要做检查的分支名
  art:
    # 此分支允许提交的文件类型，其他类型文件都不允许提交，"-"减号表示这是个列表
    - .asset
    - .fbx
```
