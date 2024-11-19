class _Bind {
    form(name: string) {
        return {
            name: name,         // 这应该是一个用于描述字段的名称
            label: name,        // 标签文本
            placeholder: `请输入${name}`, // 输入框的占位符
            rules: [{required: true, message: `请填写${name}`}] // 校验规则
        }
    }

}

export const Bind = new _Bind()