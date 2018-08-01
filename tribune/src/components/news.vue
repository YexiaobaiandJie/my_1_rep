<template>
    <div class="content">
        <div class="newtitle">
            <h2>新闻>></h2>
        </div>
        <hr />
            <div  v-for="newitem in newitems" class="newstyle">
                <div class="new-l">{{newitem.title}}</div>
                <div class="new-r">{{newitem.date}}</div>
            </div>
    </div>
</template>


<script>
export default{
    mounted:function(){
        this.loadnews()
    },
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
    },
    methods:{
        loadnews:function(){
        this.$http.get('http://localhost:8081/news?token=5b5ff5794cf88aed639018e0&pagesize=10').then(response => {
            console.log(response.data);
        },response => {
            console.log("error");
        });
        }
    }
}
</script>

<style scoped>
.content{
    width:90%;
    margin-left:5%;
    height:100%;
    
}
.newtitle{
    text-align:left;
    padding-left:10px;
    margin-bottom:-10px;
    margin-top:-2px;
   
}
.newstyle{
    width:100%;
    
}
.new-l{
   text-align: left;
   height:10%;
   width:70%;
   float: left;
   padding-left:1%;
   font-size:120%;
   padding-bottom:0.4%;
   padding-top:0.4%;
   
}
.new-r{
   text-align:right;
   height:10%;
   width:27%;
   float: left;
   font-size:120%;
   padding-right:15px;
   padding-bottom: 0.4%;
   padding-top:0.4%;
   
}
</style>