const   STORAGE_KEY='todos-vuejs'
export default{
    fetch: function(key){
        return JSON.parse(window.localStorage.getItem(key) || '[]')
    },
    save: function(items,key){
        window.localStorage.setItem(key,JSON.stringify(items))
    },
    remove:function(key){
        window.localStorage.removeItem(key)
    }
}