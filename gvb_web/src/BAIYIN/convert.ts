function empty(str: string){
    return str.replace(/ /g, "")
}

function Field(str: string) {
    if (empty(str) === "") return
    let s = str.replace(/ |\n|{|}/g, "")
    let arr = s.replace(/\[[^\]]*\]/, "[]");
    let len: string[] = arr.split(",")
    let result = "\ninterface  {\n"
    for (const i in len) {
        let r = len[i].split(":")
        let r1 = r[0].replace(/"/g, "")
        let r2 = r[1].substring(0, 1)
        let t = "number"
        switch (r2) {
            case '"':
                t = "string"
                break
            case '[':
                t = "[]"
                break
            case 't' || 'f':
                t = "boolean"
                break
        }
        let res = r1 + ":" + t + "\n"
        result += res
    }
    result += "}"
    console.log(result)
}

function Value(str: string) {
    if (empty(str) === "") return
    let s = str.replace(/ /g, "")
    let len: string[] = s.split("\n")
    for (const i in len) {
        let r = len[i].split(":")
        let r1 = r[0]   // "\"" + r[0] + "\""
        let r2: any = r[1]
        let t: any = "[]"
        switch (r2) {
            case "string":
                t = '""'
                break
            case "number":
                t = 0
                break
            case "boolean":
                t = false
                break
        }
        let res = r1 + ":" + t + ","
        console.log(res)
    }
}


Field(`       `);


Value(`         created_at: string
        id: number
        name: string
        icon: string
        abstract: string
        tags: RoleDetailTag[]
        chat_total: number  `);
