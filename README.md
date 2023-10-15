# MDAddTag
- 为你的md文件按照目录结构添加tag.
- 如果你刚从传统的目录结构的笔记软件如：OneNote, EverNote, 等等。转移到双链笔记应用如：LogSeq, 
Obsidian, 那你可能也有和我一样的问题在；在目录结构应用中辛苦维护的结构关系全部不见。在图谱中看不到关联关系， 
也不能享受方便的双链跳转功能。因此你就需要这个工具来帮助你为你的每个笔记添加所属目录的标签
## 举例
你有下面的目录结构的文件
```text
├─docker
│  │  │  ├─Dockerfile
│  │  │  │  ├─ARG
│  │  │  │  └─RUN
│  │  │  └─Nano Server

```
转换以后， 在ARG文件中的第一行会出现#Dockerfile， 如果文件夹名称有空格，则会被剔除。
---
---
# MDAddTag
- Add tags to your md files based on the directory structure.
- If you have recently transitioned from traditional note-taking software with directory structures such as OneNote, EverNote, etc. to a bidirectional note-taking app like LogSeq or Obsidian, you might have encountered the same problem as me. All the structural relationships that were painstakingly maintained in the directory structure application are lost. You can't see the connections in the graph, nor can you enjoy the convenience of bidirectional linking. That's why you need this tool to help you add tags based on the directory structure for each of your notes.
Example

Suppose you have the following directory structure for your files:
```text

├─docker
│  │  │  ├─Dockerfile
│  │  │  │  ├─ARG
│  │  │  │  └─RUN
│  │  │  └─Nano Server
```
After the conversion, the first line in the ARG file will have the tag "#Dockerfile". If a folder name has a space, it will be excluded.
