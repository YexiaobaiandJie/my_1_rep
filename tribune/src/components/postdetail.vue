<template>
    <div class="content">
        <div class="posttitle1">
            <h2 v-on:click="backtopost">论帖-author</h2>
        </div>
        <hr />

        
            <h3 class="posttitle">{{this.details.title}}</h3>
            <h3 class="postauthor">author:{{this.details.author}}</h3>
            <div class="postcontent"><p>{{this.details.content}}</p></div>
            <div class="postdate">date:{{this.localdate}}</div>
            <hr />
            <div class="comment" >
                <div>
                <div class="total_comment">{{comemntarea}}:共有{{comment_count}}条评论</div>
                <button class="comment_btn" v-bind:class="{incomment:isincomment}" v-on:click="topublishcomment">comment</button>
                </div>
                <hr />
                <div v-show="tocomment">
                    <div>
                        <textarea id="comment_text" class="comment_publish_text"></textarea>
                    </div>
                    <button class="comment_publish_btn" v-on:click="backtopostdetail();publish_comment()">publish</button>
                </div>
                <hr />
                <div v-for="comitem in comitems">
                    <div class="comment_userid">{{comitem.userid}}</div>
                    <div class="comment_content">{{comitem.com}}</div>
                    <div class="comment_date">{{comitem.time}}</div>
                    <hr />
                </div>
            </div>
    </div>
</template>


<script>
import Store from './store'
export default{
    mounted:function(){
        this.loadpostings()
    },
    data(){
        return{
            message:"this is the postdetail",
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
            comemntarea:"评论区",
            tocomment:false,
            isincomment:false,
            localdate:""
        }
    },
    methods:{
        loadpostings:function(){
            this.author=this.$route.query.Author
            this.date=this.$route.query.Date
            console.log(this.Date)
            this.$http.get('http://localhost:3000/postings/postdetail?Author='+this.author+'&Date='+this.date).then(function(response){
                this.details=response.body;
                this.comitems=this.details.com;
                this.comment_count=this.details.comcount
                this.localdate= this.timestampToTime(this.details.date)
            },function(response){
                console.log("error");
            }); 
        },
        onbtn:function(){
            this.ison=true
        },
        gobtn:function(){
            this.ison=false
        },
        backtopost:function(){
            this.isShow=true
        },
        topublishcomment:function(){
            // this.$router.push({path:'/comment',query:{Author:this.author,Date:this.date}})
            this.tocomment=true
            this.isincomment=true
        },
        backtopostdetail:function(){
            this.tocomment=false
            this.isincomment=false
        },
        publish_comment:function(){
            if(typeof(this.date)!="int"){   //如果date类型不是int64
                this.date=parseInt(this.date)
            }
            var token=Store.gettoken()
            var area=document.getElementById("comment_text")
            var comment_content=area.value
            this.$http.post('http://localhost:3000/comment',{Author:this.author,Date:this.date,Token:token,Com:comment_content})
            this.$router.go()
        },
        timestampToTime:function (timestamp)  {
            var date = new Date(timestamp * 1000);
            var Y = date.getFullYear() + '-';
            var M = (date.getMonth()+1 < 10 ? '0'+(date.getMonth()+1) : date.getMonth()+1) + '-';
            var D = date.getDate() + ' ';
            var h = date.getHours() + ':';
            var m = date.getMinutes() + ':';
            var s = date.getSeconds();
            return Y+M+D+h+m+s;
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
.total_comment{
    float: left;
}
.comment_btn{
    margin-left:20px;
    font-size:15px;
}
.comment_publish_text{
    float: left;
    width:90%;
    height:50px;
    font-size:17px;
}
.comment_publish_btn{
    margin-left:10px;
    height:55px;
    border:none;
    border-radius: 2em;
    outline:none;
}
.incomment{
    background-color:#0ff;
    color:aliceblue;
    border:none;
    outline:none;
}
</style>