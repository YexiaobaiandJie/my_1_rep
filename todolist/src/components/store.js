const   STORAGE_KEY='todos-vuejs'
export default{
    fetch: function(key){
        return JSON.parse(window.localStorage.getItem(key) || '[]')
    },
    save: function(items){
        window.localStorage.setItem(STORAGE_KEY,JSON.stringify(items))
    },
    clear:function(){
        window.localStorage.clear(STORAGE_KEY)
    }
}