# 由Go编写的文件备份程序

## 主要用途

- 用于定时备份手机QQ聊天记录

## 其他用途（需要更改保存路径和待备份文件路径）

- 备份某些文件到另一些地方

## 实现

- 使用shell进行文件备份
- 使用gocron进行定时备份

## 在Termux上执行

- 需要Root权限访问`/data/data/`下的文件
- 所有配置文件会在首次执行程序时生成在`/sdcard/Android/`下，文件名为`BackConfig.json`
- 更改配置(如更改定时时间)时程序将会在之前的所设时间执行一遍程序后才会使新设置的时间生效(其他配置同理)

## 打包为Magisk模块

- 作为Magisk模块可以更好执行

## TODO

- 将文件通过WebSocket备份到远端
- 更改配置文件生成位置
