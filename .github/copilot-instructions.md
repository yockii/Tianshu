本项目主要为go语言编写，使用Go 1.25版本。
cmd中对三端（用户端、管理端、指挥大屏端）做了启动入口的分离，分别编译出三个不同的服务程序

在fontends中是三端的前端代码，使用了Vue3+ts+vite来构建，并在各自文件夹中写了MODULES.md进行说明

# MCP Interactive Feedback 规则

1. 在任何流程、任务、对话进行时，无论是询问、回复、或完成阶段性任务，皆必须调用 MCP mcp-feedback-enhanced。
2. 每当收到用户反馈，若反馈内容非空，必须再次调用 MCP mcp-feedback-enhanced，并根据反馈内容调整行为。
3. 仅当用户明确表示「结束」或「不再需要交互」时，才可停止调用 MCP mcp-feedback-enhanced，流程才算结束。
4. 除非收到结束指令，否则所有步骤都必须重复调用 MCP mcp-feedback-enhanced。
5. 完成任务前，必须使用 MCP mcp-feedback-enhanced 工具向用户询问反馈。