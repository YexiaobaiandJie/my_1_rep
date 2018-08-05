<template>
<div class="content">
    <div class="newtitle">
        <h2 class="backnews" v-on:click="loadnews">新闻>></h2>
    </div>
        <hr />
    <div>
        <h2 class="newstitle">{{title}}</h2>
        <h3 class="newsauthor">author:{{author}}</h3>
        <div class="newscontent">{{content}}</div>
        <div class="per" v-on:click="openh">查看全文</div>
        <div class="newsdate">{{date}}</div>
    </div>     
    <hr />
    <div>
                <div>
                <div>{{comemntarea}}:共有{{comment_count}}条评论</div>
                <button>comment</button>
                </div>
                <hr />
                <div>
                    <div>
                        <textarea id="comment_text"></textarea>
                    </div>
                    <button v-on:click="publish_comment">publish</button>
                </div>
                <hr />
                <div v-for="comitem in comitems">
                    <div>{{comitem.userid}}</div>
                    <div>{{comitem.com}}</div>
                    <div>{{comitem.date}}</div>
                    <hr />
                </div>
            </div>
</div>  
</template>

<script>
import Store from './store'
export default{
    
    data(){
        return{
            message:"this is the newsdetail",
            title:"",
            author:"",
            content:"",
            date:"",
            newsurl:"",
            newsid:"12122",
            comitems:[],
            comment_count:0,
            comemntarea:"评论区"
        }
    },
    methods:{
        openh:function(){
            window.location.href=this.newsurl;
        },
        loadnews:function(){
            this.newsid=this.$route.query.id
             this.$http.get('http://localhost:3000/newsd?id='+this.newsid).then(function(res){
                var data = res.body
                console.log(data)
                this.title=data.newsinfo.title
                this.author=data.newsinfo.authorName
                this.date=data.newsinfo.publishDate
                this.content=data.newsinfo.summaryAuto
                this.newsurl=data.newsinfo.url
                this.comitems=data.newscomment
                },function(res){
                alert("返回新闻数据出错!")
            })
            
            
            // console.log(this.title)
        },
        publish_comment:function(){
            if(typeof(this.newsid)!="int"){   //如果date类型不是int
                this.newsid=parseInt(this.newsid)
            }
            var token = Store.gettoken()
            var newsid = this.newsid
            var area=document.getElementById("comment_text")
            var comment_content=area.value
            this.$http.post('http://localhost:3000/news/comment',{Newsid:newsid,Token:token,Com:comment_content})
            
            console.log("评论函数运行完毕")
            this.$router.go()
        }
    },
    mounted:function(){
        this.loadnews()
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
.newstitle{
    text-align:center;
}
.newsauthor{
    text-align: right;
    padding-right:8%;
}
.newscontent{
    width:80%;
    height:100%;
    margin-left:10%;
    font-size:18px;
}
.newsdate{
    padding-right:10%;
    width:50%;
    float:right;
    text-align: right;
}
.per{
    color:red;
    margin-left:10%;
    cursor: pointer;
    width:8%;
    float:left;
}
</style>