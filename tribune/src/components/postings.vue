<template>
    <div class="content">
        <div class="posttitle1">
            <h2 v-on:click="backtopost">论帖-author</h2>
        </div>
        <hr />
        <div v-show="isShow">
            <div  v-for="postitem in postitems" class="poststyle">
                <div class="post-l" v-on:click="clickpost(postitem)">{{postitem.title}}</div>
                <div class="post-r">{{postitem.author}}</div>
            </div>
            <button class="btnpublish" v-bind:class="{clickbtn:ison}" v-on:mouseover="onbtn" v-on:mouseout="gobtn">发帖</button>
        </div>

        <div v-show="!isShow">
            <h3 class="posttitle">{{this.details.title}}</h3>
            <h3 class="postauthor">author:{{this.details.author}}</h3>
            <div class="postcontent"><p>{{this.details.content}}</p></div>
            <div class="postdate">date:{{this.details.date}}</div>
            <hr />
            <div class="comment" >
                <div>{{comemntarea}}:共有{{comment_count}}条评论</div>
                <hr />
                <div v-for="comitem in comitems">
                    <div class="comment_userid">{{comitem.userid}}</div>
                    <div class="comment_content">{{comitem.com}}</div>
                    <div class="comment_date">{{comitem.time}}</div>
                    <hr />
                </div>
            </div>
        </div>
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
            this.isShow=false
            this.$http.post('http://localhost:3000/detail',{"Author":postitem.author,"Date":postitem.date}).then(response => {
                console.log(response.data)
                this.details=response.data;
                this.comitems=this.details.com;
                this.comment_count=this.details.comcount
            },response => {
                console.log("error");
            });  
        },
        backtopost:function(){
            this.isShow=true
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