export namespace DefineUtils {
    export type uid = { id: number; };
    export type options = { label: string; value: string | number };
    export type optionsList = { label: string; value: string | number; search?: string; option: options[]; };

    export const Role = [
        {label: "管理员", value: "管理员"},
        {label: "普通用户", value: "普通用户"},
        {label: "游客", value: "游客"},
        {label: "禁用用户", value: "禁用用户"},
    ];
}