interface jwtInfoType {
    role: number
    user_id: number
    exp: number
    iss: string
}


class _JWT {
    constructor() {
    }

    Parse(token: string): jwtInfoType {
        let payload = token.split(".")[1]
        return JSON.parse(decodeURIComponent(escape(
            window.atob(payload.replace(/-/g, "+").replace(/_/g, "/")))))
    }
}

export const JWT = new _JWT();

