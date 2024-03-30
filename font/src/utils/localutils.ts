
export const saveLocalUser = (token:string,userid:string,useraccount:string,usernickname:string) =>{
    localStorage.setItem("token",token)
    localStorage.setItem("userid",userid)
    localStorage.setItem("user_nickname",usernickname)
    localStorage.setItem("user_account",useraccount)
}

export const deleteLocalUser = () =>{
    localStorage.removeItem("token")
    localStorage.removeItem("userid")
    localStorage.removeItem("user_account")
}
