@echo off
::The current program needs to be executed.
::the CMD directory under the installation path needs to be added to the system environment Path variable.

::echo .....Jump to folder E:\komi\komiTest（跳转至该路径，是一个git项目，应该会有一个.git文件夹。如果这个bat文件已经在git项目路径下了，可以不用跳转，直接写git命令，会对该git项目直接进行操作）
::cd /d E:\komi\komiTest

::下面说一下push.default matching和push.default simple的区别：
::push.default设置maching的意思是：git push 会把你本地所有分支push到名称相对应的远程主机上。这意味着可能你会在不经意间push一些你原本没打算push的分支。
::push.default设置成simple的意思是：git push仅仅把当前所在分支push到从当初git pull pull下来的那个对应分支上，另外，这个过程也会同时检查各个分支的名称是否相对应。
 
::echo .....!!!!git checkout master （执行git命令：切换至master分支）
::git checkout master
 
git status
git add -A --all
git commit -m "bat auto commit log"
::git push origin master
git push.default simple

echo ///////////////////////task done

echo .....enter any key to exit（想要bat文件执行完可以看到日志，而不是自动关闭，用以下方式可以实现：pause>nul）
pause