# Capistrano

[github](https://github.com/capistrano/capistrano)

## Capistrano

A deployment automation tool built on Ruby, Rake and SSH.

Capistrano is a framework for building automated deployment script. Although Capistrano itself is written in Ruby, it can easily be used to deploy project of any language or framework, be it Rails, Java, or PHP.

Once installed, Capistrano gives you a ```cap``` tool to perform your deployments from the comfort of your command line.

```
cd xxxxx_project
cap production deploy
```

When you run ```cap```, Capistrano dutifully connects to your server(s) via SSH and executes the step necessary to deploy your project. You can define those steps yourself by writing Rake tasks, or by using pre-built task libraries provided by the Capistrano community.