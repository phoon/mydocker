# mydocker

《自己动手写Docker》一书的复刻，使用Go modules(Go 1.11+)构建依赖。

在原书代码基础上更改：

- 添加了对USER_NAMESPACE的支持，Archlinux内核虽然开启了对USER_NAMESPACE的支持，但默认只限于root用户，使用`sudo sysctl kernel.unprivileged_userns_clone=1 `即可允许普通用户运行无特权容器。
- 在启动容器前，使用`sudo mount --make-private -t proc proc /proc -o remount` ，可避免容器进程对宿主机/proc挂载的破坏。

