const   STORAGE_KEY='todos-vuejs'
const   NOTE_KEY='note_key'
export default{
    fetch: function(key){
        return JSON.parse(window.localStorage.getItem(key) || '[]')
    },
    fetchnote: function(){
        return JSON.parse(window.localStorage.getItem(NOTE_KEY) || '[]')
    },
    savenote: function(noteitems){
        window.localStorage.setItem(NOTE_KEY,JSON.stringify(noteitems))
    },
    save: function(items,key){
        window.localStorage.setItem(key,JSON.stringify(items))
    },
    remove:function(key){
        window.localStorage.removeItem(key)
    },
    removenote:function(thisnote){
        var note_key=JSON.parse(window.localStorage.getItem(NOTE_KEY))
        for(var i=0;i<note_key.length;i++)
        {
            if(note_key[i].label===thisnote){
                note_key.splice(i,1)
            }
        }
        window.localStorage.setItem(NOTE_KEY,JSON.stringify(note_key))
    }
}