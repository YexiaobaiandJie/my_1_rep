const STORAGE_KEY='token'

export default{
    savetoken :function(token){
        window.localStorage.setItem(STORAGE_KEY,JSON.stringify(token))
    },
    gettoken :function(){
        return JSON.parse(STORAGE_KEY)
    } 
}