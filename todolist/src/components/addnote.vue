<template>
    <div class="logtext1">
        <div style="line-heigh:82px;font-size:28px;margin-top:10px;">NEW NOTE</div>
        <br />
        
        <div class="logtext1d">
            <span>name</span>
            <input class="input1" type="text" v-model="evname"/>
        </div>
        <div class="logtext2d">
            <span>time</span>
            <input class="input2" type="date" value="" v-model="statime"/>
        </div>
        <br />
        <button class="btncr" v-on:click="cre_note">create</button>
        <button class="btncr" v-on:click="closelog">cancel</button>
        <br />
        <br />
    </div> 
</template>

<script>
export default {
    props:{
        parnoteitems:{
            type:Array,
            required:true
        }
    },
    data(){
       return{
            evname:'',
            statime:'2018-07-20',
            noteinfo:
                {
                    label:'',
                    time:''
                }
       }
   },
   methods:{
       cre_note:function(){
           var count=0
           for(var i=0;i<this.parnoteitems.length;i++){
               if(this.evname===this.parnoteitems[i].label){
                   count++;
                   break;
               }
           }
           if(count===0){
                this.noteinfo.label=this.evname
                this.noteinfo.time=this.statime
                this.$emit('getNoteInfo', this.noteinfo)
                this.evname=''
                this.statime='2018-07-20'
           }else{
               this.evname=''
               this.statime='2017-07-20'
               alert("name already exists")
           }
        },
       closelog:function(){
           this.$parent.ChangeCreatenoteStatue()
       }
   }
    
}
</script>

<style>
    .logtext1{
        max-width:270px;
        max-height:200px;
        min-width:250px;
        min-height:190px;
        width:27%;
        height:40%;
        top: 50%;   
        left: 50%;   
        -webkit-transform: translate(-50%, -50%);   
        -moz-transform: translate(-50%, -50%);   
        -ms-transform: translate(-50%, -50%);   
        -o-transform: translate(-50%, -50%);   
        transform: translate(-50%, -50%);  
        background-color:#6b6260;
        position: absolute;
        z-index:10;
        border-radius: 0.5em;
        color:lemonchiffon;
        
    }
    .logtext1d{
        font-size:15px;
        margin-left:-10%;
        margin-right:10px;
    }
    .logtext2d{
        font-size:15px;
        margin-top:15px;
        margin-left:-11%;
    }
    .input1{
        font-size:17px;
        width:52%;
        height:23px;
    }
    .input2{
        
        height:23px;
        font-size:17px;
    }
    .btncr{
        width:90px;
        height:30px;
        font-size:20px;
        color:lavender;
        background-color: #6b6260;
        border:0.5px solid;
    }
    
</style>