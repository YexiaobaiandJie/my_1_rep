<template>
    <div class="content">
        <div class="posttitle1">
            <h2 v-on:click="backtopost">论帖-author</h2>
        </div>
        <hr />
            <div  v-for="postitem in postitems" class="poststyle">
                <div class="post-l" v-on:click="clickpost(postitem)">{{postitem.title}}</div>
                <div class="post-r">{{postitem.author}}</div>
            </div>
            <button class="btnpublish" v-on:click="topublish" v-bind:class="{clickbtn:ison}" v-on:mouseover="onbtn" v-on:mouseout="gobtn">发帖</button>
        
    </div>
</template>


<script>

export default{
    mounted:function(){
        this.loadpostings()
    },
    data(){
        return{
            title:"title",
            author:"author",
            content:"content",
            date:"date",
            ison:false,
            postitems:[],
            comitems:[],
            details:"",
            isShow:true,
            comment_count:0,
            comemntarea:"评论区"
        }
    },
    methods:{
        loadpostings:function(){
            this.$http.get('http://localhost:3000/postings').then(response => {
                this.postitems=response.data;
            },response => {
                console.log("error");
            });
        },
        onbtn:function(){
            this.ison=true
        },
        gobtn:function(){
            this.ison=false
        },
        clickpost:function(postitem){
            this.$router.push({path:'/postings/postdetail',query:{Author:postitem.author,Date:postitem.date}})
        },
        backtopost:function(){
            this.isShow=true
        },
        topublish:function(){
            this.$router.push({path:'/postings/publish'})
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
/* .posttitle{
    text-align:left;
    padding-left:10px;
    margin-bottom:-10px;
    margin-top:-2px;
    
} */
.posttitle1{
    text-align:left;
    padding-left:10px;
    margin-bottom:-10px;
    margin-top:-2px;
    cursor: pointer;
}
.poststyle{
    width:100%;
    
}
.post-l{
   text-align: left;
   height:10%;
   width:70%;
   float: left;
   padding-left:1%;
   font-size:120%;
   padding-bottom:0.4%;
   padding-top:0.4%;
   cursor: pointer;
}
.post-r{
   text-align:right;
   height:10%;
   width:27%;
   float: left;
   font-size:120%;
   padding-right:15px;
   padding-bottom: 0.4%;
   padding-top:0.4%;
}
.btnpublish{
    width:8%;
    font-size:108%;
    padding-top:4px;
    padding-bottom: 4px;
    margin-top:2%;
    margin-right:1%;
    float:right;
    background-color:#555;
    border:none;
    color:aliceblue;
    transition:0.7s all;
    outline:none;
}
.clickbtn{
    background-color:#ff4500;
    transition:0.7s all;
} 
.posttitle{
    text-align:center;
}
.postauthor{
    text-align: right;
    padding-right:8%;
}
.postcontent{
    width:80%;
    height:100%;
    margin-left:10%;
    font-size:18px;
}
.postdate{
    text-align: right;
    padding-right:8%;
}
.comment{
    width:80%;
    height:100%;
    margin-left:10%;
}
.comment_date{
    text-align:right;
    padding-right:1%;
}
.comment_userid{
    text-align:left;
    padding-left:1%;
}
.comment_content{
    padding-left:2%;
    padding-right:2%;
}
</style>