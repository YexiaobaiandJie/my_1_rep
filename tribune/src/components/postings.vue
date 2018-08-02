<template>
    <div class="content">
        <div class="posttitle">
            <h2>论帖~~author</h2>
        </div>
        <hr />
            <div  v-for="postitem in postitems" class="poststyle">
                <div class="post-l">{{postitem.title}}</div>
                <div class="post-r">{{postitem.author}}</div>
            </div>
            <button class="btnpublish" v-bind:class="{clickbtn:ison}" v-on:mouseover="onbtn" v-on:mouseout="gobtn">发帖</button>
    </div>
</template>


<script>
export default{
    mounted:function(){
        this.loadpostings()
    },
    data(){
        return{
            ison:false,
            postitems:[]
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
.posttitle{
    text-align:left;
    padding-left:10px;
    margin-bottom:-10px;
    margin-top:-2px;
   
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

</style>