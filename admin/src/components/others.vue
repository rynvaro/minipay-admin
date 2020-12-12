<template>
  <div>
    <el-row>
      <el-col :span="3">
        <div class="title">Home背景图片：</div>
      </el-col>
      <el-col :span="4">
         <el-upload class="upload-demo" ref="item" action="http://localhost:8090/upload" :on-success="onItemUploadSuccess" :auto-upload="false">
           <el-button slot="trigger" size="small" type="primary">选取文件</el-button>
           <el-button style="margin-left: 10px;" size="small" type="success" @click="submitUpload(1)">上传到服务器</el-button>
           <div slot="tip" class="el-upload__tip"></div>
         </el-upload>
      </el-col>
      <el-col :span="4">
        <el-button style="width: 100%;" @click="update">更新</el-button>
      </el-col>
    </el-row>

    <el-row style="margin-top: 20px;">
      <el-col :span="12">抽奖资格审核</el-col>
    </el-row>
    <el-row>
      <el-col :span="10">
        <el-input placeholder="输入手机号,然后回车" v-model="phone" @change="checkLottory" clearable></el-input>
      </el-col>
      <el-col :span="4">
        <el-button style="width: 100%;" @click="checkLottory">审核</el-button>
      </el-col>
    </el-row>

    <el-divider></el-divider>

    <el-divider content-position="left">返现配置</el-divider>
    <el-row>
      <el-col :span="3" style="margin-top: 10px">
        最高返现： 
      </el-col>
      <el-col :span="4">
        <el-input placeholder="最高返现(元)" v-model.number="fanxian.maxFanxian" clearable></el-input>
      </el-col>
    </el-row>
    <el-row style="margin-top: 15px">
      <el-col :span="3" style="margin-top: 10px">
        首次返现百分比： 
      </el-col>
      <el-col :span="4">
        <el-input placeholder="如：40" v-model.number="fanxian.firstFanxianPercent" clearable></el-input>
      </el-col>

      <el-col :span="3" style="margin-top: 10px">
        二次返现百分比： 
      </el-col>
      <el-col :span="4">
        <el-input placeholder="如：30" v-model.number="fanxian.secondFanxianPercent" clearable></el-input>
      </el-col>

      <el-col :span="3" style="margin-top: 10px">
        三次返现百分比： 
      </el-col>
      <el-col :span="4">
        <el-input placeholder="如：20" v-model.number="fanxian.thirdFanxianPercent" clearable></el-input>
      </el-col>
    </el-row>
    <el-row>
      <el-col :span="24" style="margin-top: 15px">
        <el-button type="primary" @click="updateFanxian">提交</el-button>
      </el-col>
    </el-row>
    <el-divider></el-divider>

  </div>
</template>

<script>
  import axios from 'axios'
  import { Loading } from 'element-ui';
  export default {
    data() {
      return {
        image: '',
        phone: '',
        fanxian: {},
      }
    },
    methods: {
      checkLottory() {
        if (!this.phone) {
          this.$message.error("请输入手机号")
          return
        }
        let data = {
          tp: 'checklottory',
          phone: this.phone,
        }
        let loadingInstance = Loading.service({ fullscreen: true ,text:'审核中...'});
        axios.post('http://localhost:8090/proxy',data).then(resp => {
          loadingInstance.close()
          if (resp.status == 200) {
            if (resp.data == -1) {
              this.$message.error("用户不存在")
              return
            }
            this.$message.success("通过审核")
          }else {
            this.$message.error("审核失败")
          }
        })
      },
      updateFanxian() {
        let loadingInstance = Loading.service({ fullscreen: true ,text:'更新中...'});
        let data = this.fanxian
        data.tp = "fanxian"
        axios.post('http://localhost:8090/proxy',data).then(resp => {
          if (resp.status == 200) {
            this.$message.success("更新成功")
          }else {
            this.$message.error("更新失败")
          }
          loadingInstance.close()
        })
      },
      update(){
        if (!this.image) {
          this.$message.error("请上传图片")
          return
        }
        let data = {
          tp: 'homeheader',
          image: this.image
        }
        axios.post('http://localhost:8090/proxy',data).then(resp => {
          if (resp.status == 200) {
            this.$message.success("更新成功")
          }else {
            this.$message.error("更新失败")
          }
        })
      },
      onItemUploadSuccess(e){
        this.image = e
      },
      submitUpload(i){
        if (i == 1) {
          this.$refs.item.submit()
        }else if (i == 2) {
          this.$refs.itemDetail.submit()
        }

      }
    },
    mounted() {
      axios.post('http://localhost:8090/proxy',{tp: "fanxianconfig"}).then(resp => {
          if (resp.status == 200) {
            console.log(resp.data)
            this.fanxian = resp.data.data
          }
        })
    }
  }
</script>

<style>
  .title{
    margin-top: 10px;
  }
</style>
