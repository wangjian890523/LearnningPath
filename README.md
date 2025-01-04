#LearnningPath
之前学习的内容整理

1. 本地
    git init
2. github 创建仓库

3.首次提交 
    echo "#LearnningPath" >> README.md
    git add README.md
    git commit -m "first commit."
    git branchh -M main
    git remote  add  origin git@github.com:wangjian890523/LearnningPath.git
    git push -u orgin main 

4. ssh 问题
    1. 生成新的ssh key 添加到github
    2. 本地know_hosts github 变更更新
        ssh-keygen -R github.com
    3. 再次ssh
        ssh git@github.com
    4. 测试
        ssh -T git#@github.com
5. 其它问题
    添加remote url 错误，github.com 后使用了"/"
    修改
    git remote set-url origin  ssh_url




