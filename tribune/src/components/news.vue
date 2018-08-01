<template>
    <div class="content">
        <div class="newtitle">
            <h3>新闻>></h3>
        </div>
        <hr />

        <ul>
            <li  v-for="newitem in newitems" class="newitemstyle">{{newitem.title}}</li>
        </ul>
    </div>
</template>


<script>
export default{
    data(){
        return{
            message:"this is the news component",
            newitems:[
                {
                    title:"新闻的题目",
                    date:"2018-08-01"
                },
                {
                    title:"新闻的题目",
                    date:"2018-08-01"
                },
                {
                    title:"新闻的题目",
                    date:"2018-08-01"
                },
                {
                    title:"新闻的题目",
                    date:"2018-08-01"
                }
            ]
        }
    }
}
</script>

<style scoped>
.content{
    width:90%;
    margin-left:5%;
    height:100%;
    background:#ff4500;
}
.newtitle{
    text-align:left;
    padding-left:10px;
    margin-bottom:-10px;
}
.newitemstyle{
    text-align: left;
}
</style>