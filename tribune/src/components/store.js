const TOKEN_KEY='token'
const USERID_KEY='userid'
const STATUS_KEY='islogin'

export default{
    savetoken :function(token){
        window.localStorage.setItem(TOKEN_KEY,token)
    },
    gettoken :function (){
        return window.localStorage.getItem(TOKEN_KEY)
    },
    saveuseid :function(userid){
        window.localStorage.setItem(USERID_KEY,userid)
    },
    getuserid :function (){
        return window.localStorage.getItem(USERID_KEY)
    },
    setlogin:function(status){
        window.localStorage.setItem(STATUS_KEY,JSON.stringify(status))
    },
    getstatus:function(){
        return window.localStorage.getItem(STATUS_KEY)
    }

}