<template>

  <div>
    <el-row>
      <el-col :span="2">
        <div class="title">店铺名称：</div>
      </el-col>
      <el-col :span="12">
        <el-input placeholder="请输入店铺名称" v-model="storeName" clearable></el-input>
      </el-col>
    </el-row>

    <el-row>
      <el-col :span="2">
        <div class="title">商家描述：</div>
      </el-col>
      <el-col :span="12">
        <el-input type="textarea" placeholder="请输入商家描述" v-model="storeDesc" clearable></el-input>
      </el-col>
    </el-row>

    <el-row>
      <el-col :span="2">
        <div class="title">营业时间：</div>
      </el-col>
      <el-col :span="12">
        <el-time-select placeholder="起始时间" v-model="startTime" :picker-options="{
              start: '00:00',
              step: '00:30',
              end: '24:00'
            }">
        </el-time-select>
        <el-time-select placeholder="结束时间" v-model="endTime" :picker-options="{
              start: '00:00',
              step: '00:30',
              end: '24:00',
              minTime: startTime
            }">
        </el-time-select>
      </el-col>
    </el-row>

    <el-row>
      <el-col :span="2">
        <div class="title">折扣信息：</div>
      </el-col>

      <el-col :span="4">
        <el-select v-model="discount" clearable placeholder="选择折扣">
            <el-option
              v-for="item in options"
              :key="item.value"
              :label="item.label"
              :value="item.value">
            </el-option>
          </el-select>
      </el-col>

      <el-col :span="8">
        <el-date-picker
          v-model="discountDate"
          type="daterange"
          range-separator="至"
          start-placeholder="开始日期"
          end-placeholder="结束日期"
          value-format="timestamp">
        </el-date-picker>
      </el-col>
    </el-row>

    <el-row>
      <el-col :span="2">
        <div class="title">店铺类型：</div>
      </el-col>
      <el-col :span="12">
        <el-radio-group v-model="storeType">
          <el-radio-button label="1">餐饮</el-radio-button>
          <el-radio-button label="2">娱乐</el-radio-button>
          <el-radio-button label="3">其他</el-radio-button>
        </el-radio-group>
      </el-col>
    </el-row>

    <el-row>
      <el-col :span="2">
        <div class="title">轮播视频：</div>
      </el-col>
      <el-col :span="4">
        <el-upload class="upload-demo" ref="video" action="http://localhost:8090/upload" :multiple="true" :limit="6"
          :file-list="videoList" :on-success="onUploadVideoSuccess" :auto-upload="false" accept=".mp4">
          <el-button slot="trigger" size="small" type="primary">选取文件</el-button>
          <el-button style="margin-left: 10px;" size="small" type="success" @click="submitUpload(1)">上传到服务器</el-button>
          <div slot="tip" class="el-upload__tip"></div>
        </el-upload>
      </el-col>

      <el-col :span="2">
        <div class="title">轮播图片：</div>
      </el-col>
      <el-col :span="4">
        <el-upload class="upload-demo" ref="image" action="http://localhost:8090/upload" :multiple="true" :limit="6"
          :file-list="imageList" :on-success="onUploadImageSuccess" :auto-upload="false">
          <el-button slot="trigger" size="small" type="primary">选取文件</el-button>
          <el-button style="margin-left: 10px;" size="small" type="success" @click="submitUpload(2)">上传到服务器</el-button>
          <div slot="tip" class="el-upload__tip"></div>
        </el-upload>
      </el-col>
    </el-row>

    <el-row>
      <el-col :span="2">
        <div class="title">店铺门头：</div>
      </el-col>
      <el-col :span="4">
        <el-upload class="upload-demo" ref="storeImage" action="http://localhost:8090/upload" :multiple="true" :limit="6"
          :file-list="storeImageList" :on-success="onUploadStoreImageSuccess" :auto-upload="false">
          <el-button slot="trigger" size="small" type="primary">选取文件</el-button>
          <el-button style="margin-left: 10px;" size="small" type="success" @click="submitUpload(3)">上传到服务器</el-button>
          <div slot="tip" class="el-upload__tip"></div>
        </el-upload>
      </el-col>

      <el-col :span="2">
        <div class="title">产品图片：</div>
      </el-col>
      <el-col :span="4">
        <el-upload class="upload-demo" ref="productImage" action="http://localhost:8090/upload" :multiple="true" :limit="6"
          :file-list="productImageList" :on-success="onUploadProductImageSuccess" :auto-upload="false">
          <el-button slot="trigger" size="small" type="primary">选取文件</el-button>
          <el-button style="margin-left: 10px;" size="small" type="success" @click="submitUpload(4)">上传到服务器</el-button>
          <div slot="tip" class="el-upload__tip"></div>
        </el-upload>
      </el-col>
    </el-row>

    <el-row>
      <el-col :span="2">
        <div class="title">推广图片：</div>
      </el-col>
      <el-col :span="4">
        <el-upload class="upload-demo" ref="promoteImage" action="http://localhost:8090/upload" :multiple="true" :limit="6"
          :file-list="promoteImageList" :on-success="onUploadPromoteImageSuccess" :auto-upload="false">
          <el-button slot="trigger" size="small" type="primary">选取文件</el-button>
          <el-button style="margin-left: 10px;" size="small" type="success" @click="submitUpload(5)">上传到服务器</el-button>
          <div slot="tip" class="el-upload__tip"></div>
        </el-upload>
      </el-col>
    </el-row>


    <el-row>
      <el-col :span="4" v-for="(file, index) in tmpFiles" :key="index">
        <video
          v-if="file.isVideo"
          controls
          :src="file.download_url"
          style="width: 100px; height: 100px;"
           />
        <el-image
        v-if="!file.isVideo"
        :src="file.download_url"
        :fit="cover"
        style="width: 100px; height: 100px;"></el-image>
        <el-button @click="deleteImage(file.fileid, index)" type="text" size="small">删除</el-button>
      </el-col>
    </el-row>

    <el-row>
      <el-col :span="14">
        <div class="title">位置信息</div>
      </el-col>
    </el-row>

    <el-row :gutter="5">

      <el-col :span="2">
        <div class="title">经度：</div>
      </el-col>
      <el-col :span="2">
        <el-input placeholder="0.00" v-model.number="longitude" clearable></el-input>
      </el-col>

      <el-col :span="2">
        <div class="title">纬度：</div>
      </el-col>
      <el-col :span="2">
        <el-input placeholder="0.00" v-model.number="latitude" clearable></el-input>
      </el-col>

      <el-col :span="2">
        <el-button type="primary" icon="el-icon-search" @click="mapVisible=!mapVisible">地图获取</el-button>
      </el-col>
    </el-row>

    <el-row>
      <el-col :span="2">
        <div class="title">详细地址：</div>
      </el-col>
      <el-col :span="12">
        <el-input placeholder="请输入详细地址" v-model="address" clearable></el-input>
      </el-col>
    </el-row>

    <el-row :span="14" v-show="mapVisible">
      <el-input placeholder="搜索位置" v-model="keyword" clearable></el-input>
      <baidu-map>
        <bm-view :center="center" :zoom="zoom" class="map"></bm-view>
        <bm-local-search :keyword="keyword" :auto-viewport="true" @infohtmlset="infohtmlset"></bm-local-search>
      </baidu-map>
    </el-row>

    <el-row>
      <el-col :span="14">
        <div class="title">商家信息</div>
      </el-col>
    </el-row>

    <el-row>
      <el-col :span="2">
        <div class="title">商家姓名：</div>
      </el-col>
      <el-col :span="4">
        <el-input placeholder="请输入姓名" v-model="merchantName" clearable></el-input>
      </el-col>

      <el-col :span="2">
        <div class="title">手机号：</div>
      </el-col>
      <el-col :span="4">
        <el-input placeholder="请输入手机号" v-model="merchantPhone" clearable></el-input>
      </el-col>

    </el-row>

    <el-row>
      <el-col :span="2">
        <div class="title">银行卡号：</div>
      </el-col>
      <el-col :span="12">
        <el-input placeholder="请输入内容" v-model="merchantBankCard" clearable></el-input>
      </el-col>
    </el-row>

    <el-row>
      <el-col :span="14">
        <el-button type="primary" plain @click="doPublish">发布商家</el-button>
      </el-col>
    </el-row>
  </div>
</template>

<script>
  import axios from 'axios'
  export default {
    data() {
      return {
        id: this.$route.params._id,

        tmpFiles: [],

        mapVisible: false,

        storeName: '',
        storeDesc: '',
        startTime: '',
        endTime: '',

        options: [
          {value: 0.0, label: '0.0 折'},
          {value: 0.1, label: '0.1 折'},
          {value: 0.2, label: '0.2 折'},
          {value: 0.3, label: '0.3 折'},
          {value: 0.4, label: '0.4 折'},
          {value: 0.5, label: '0.5 折'},
          {value: 0.6, label: '0.6 折'},
          {value: 0.7, label: '0.7 折'},
          {value: 0.8, label: '0.8 折'},
          {value: 0.9, label: '0.9 折'},
          {value: 1.0, label: '1.0 折'},
          {value: 1.1, label: '1.1 折'},
          {value: 1.2, label: '1.2 折'},
          {value: 1.3, label: '1.3 折'},
          {value: 1.4, label: '1.4 折'},
          {value: 1.5, label: '1.5 折'},
          {value: 1.6, label: '1.6 折'},
          {value: 1.7, label: '1.7 折'},
          {value: 1.8, label: '1.8 折'},
          {value: 1.9, label: '1.9 折'},
          {value: 2.0, label: '2.0 折'},
          {value: 2.1, label: '2.1 折'},
          {value: 2.2, label: '2.2 折'},
          {value: 2.3, label: '2.3 折'},
          {value: 2.4, label: '2.4 折'},
          {value: 2.5, label: '2.5 折'},
          {value: 2.6, label: '2.6 折'},
          {value: 2.7, label: '2.7 折'},
          {value: 2.8, label: '2.8 折'},
          {value: 2.9, label: '2.9 折'},
          {value: 3.0, label: '3.0 折'},
          {value: 3.1, label: '3.1 折'},
          {value: 3.2, label: '3.2 折'},
          {value: 3.3, label: '3.3 折'},
          {value: 3.4, label: '3.4 折'},
          {value: 3.5, label: '3.5 折'},
          {value: 3.6, label: '3.6 折'},
          {value: 3.7, label: '3.7 折'},
          {value: 3.8, label: '3.8 折'},
          {value: 3.9, label: '3.9 折'},
          {value: 4.0, label: '4.0 折'},
          {value: 4.1, label: '4.1 折'},
          {value: 4.2, label: '4.2 折'},
          {value: 4.3, label: '4.3 折'},
          {value: 4.4, label: '4.4 折'},
          {value: 4.5, label: '4.5 折'},
          {value: 4.6, label: '4.6 折'},
          {value: 4.7, label: '4.7 折'},
          {value: 4.8, label: '4.8 折'},
          {value: 4.9, label: '4.9 折'},
          {value: 5.0, label: '5.0 折'},
          {value: 5.1, label: '5.1 折'},
          {value: 5.2, label: '5.2 折'},
          {value: 5.3, label: '5.3 折'},
          {value: 5.4, label: '5.4 折'},
          {value: 5.5, label: '5.5 折'},
          {value: 5.6, label: '5.6 折'},
          {value: 5.7, label: '5.7 折'},
          {value: 5.8, label: '5.8 折'},
          {value: 5.9, label: '5.9 折'},
          {value: 6.0, label: '6.0 折'},
          {value: 6.1, label: '6.1 折'},
          {value: 6.2, label: '6.2 折'},
          {value: 6.3, label: '6.3 折'},
          {value: 6.4, label: '6.4 折'},
          {value: 6.5, label: '6.5 折'},
          {value: 6.6, label: '6.6 折'},
          {value: 6.7, label: '6.7 折'},
          {value: 6.8, label: '6.8 折'},
          {value: 6.9, label: '6.9 折'},
          {value: 7.0, label: '7.0 折'},
          {value: 7.1, label: '7.1 折'},
          {value: 7.2, label: '7.2 折'},
          {value: 7.3, label: '7.3 折'},
          {value: 7.4, label: '7.4 折'},
          {value: 7.5, label: '7.5 折'},
          {value: 7.6, label: '7.6 折'},
          {value: 7.7, label: '7.7 折'},
          {value: 7.8, label: '7.8 折'},
          {value: 7.9, label: '7.9 折'},
          {value: 8.0, label: '8.0 折'},
          {value: 8.1, label: '8.1 折'},
          {value: 8.2, label: '8.2 折'},
          {value: 8.3, label: '8.3 折'},
          {value: 8.4, label: '8.4 折'},
          {value: 8.5, label: '8.5 折'},
          {value: 8.6, label: '8.6 折'},
          {value: 8.7, label: '8.7 折'},
          {value: 8.8, label: '8.8 折'},
          {value: 8.9, label: '8.9 折'},
          {value: 9.0, label: '9.0 折'},
          {value: 9.1, label: '9.1 折'},
          {value: 9.2, label: '9.2 折'},
          {value: 9.3, label: '9.3 折'},
          {value: 9.4, label: '9.4 折'},
          {value: 9.5, label: '9.5 折'},
          {value: 9.6, label: '9.6 折'},
          {value: 9.7, label: '9.7 折'},
          {value: 9.8, label: '9.8 折'},
          {value: 9.9, label: '9.9 折'},
        ],
        discount: 9,
        discountDate: [],

        storeType: 1,
        longitude: '',
        latitude: '',
        address: '',

        center: {
          lng: 0,
          lat: 0
        },
        zoom: 15,
        keyword: '7hao',

        videoList: [],
        imageList: [],
        storeImageList: [],
        productImageList: [],
        promoteImageList: [],

        banners: [],
        storeImages: [],
        productImages: [],
        promoteImages: [],

        merchantName: '',
        merchantPhone: '',
        merchantBankCard: '',
      }
    },
    methods: {
      doPublish() {
        var msg = ""
        if (!this.storeName) {
          this.$message.error(msg + "请输入店铺名称")
          return
        }

        if (!this.storeDesc) {
          this.$message.error(msg + "请输入店铺描述")
          return
        }

        if (!this.startTime || !this.endTime) {
          this.$message.error(msg + "请输入营业时间")
          return
        }

        if (!this.discount) {
          this.$message.error(msg + "请选择折扣信息")
          return
        }

        if (this.discountDate.length != 2) {
          this.$message.error(msg + "请选择折扣时间")
          return
        }

        if (!this.storeType) {
          this.$message.error(msg + "请选择餐饮类型")
          return
        }

        if (this.banners.length == 0) {
          this.$message.error(msg + "请上传至少一个轮播视频或者图片")
          return
        }

        if (this.storeImages.length == 0) {
          this.$message.error(msg + "请上传门头图片")
          return
        }

        if (this.productImages.length == 0) {
          this.$message.error(msg + "请上传产品图片")
          return
        }

        if (this.promoteImages.length == 0) {
          this.$message.error(msg + "请上传推广图片")
          return
        }

        if (!this.longitude || !this.latitude) {
          this.$message.error(msg + "请输入经纬度信息")
          return
        }

        if (!this.address) {
          this.$message.error(msg + "请输入详细地址")
          return
        }

        if (!this.merchantName) {
          this.$message.error(msg + "请输入商家姓名")
          return
        }

        if (!this.merchantPhone) {
          this.$message.error(msg + "请输入商家手机号")
          return
        }

        if (!this.merchantBankCard) {
          this.$message.error(msg + "请输入商家银行卡号")
          return
        }

        // TODO url
        let data = {
          "id": this.id,
          "storeName": this.storeName,
          "storeDesc": this.storeDesc,
          "startTime": this.startTime,
          "endTime": this.endTime,
          "discount": this.discount,
          "discountStart": this.discountDate[0],
          "discountEnd": this.discountDate[1],

          "storeType": this.storeType,
          "longitude": this.longitude,
          "latitude": this.latitude,
          "address": this.address,

          "banners": this.banners,
          "storeImages": this.storeImages,
          "productImages": this.productImages,
          "promoteImages": this.promoteImages,

          "merchantName": this.merchantName,
          "merchantPhone": this.merchantPhone,
          "merchantBankCard": this.merchantBankCard,
        }
        axios.post('http://localhost:8090/publish', data)
          .then(resp => (
            // this.$router.go(0),
            this.$alert('发布成功')
          ))

      },
      onUploadVideoSuccess(e) {
        this.banners.push({isVideo: true, url: e})
        this.$message.success("上传视频成功")
      },
      onUploadImageSuccess(e) {
        this.banners.push({isVideo: false, url: e})
        this.$message.success("上传图片成功")
      },
      onUploadStoreImageSuccess(e) {
        this.storeImages.push(e)
        this.$message.success("上传店铺门头成功")
      },
      onUploadProductImageSuccess(e) {
        this.productImages.push(e)
        this.$message.success("上传产品图片成功")
      },
      onUploadPromoteImageSuccess(e) {
        this.promoteImages.push(e)
        this.$message.success("上传产品图片成功")
      },

      deleteImage(fileid, index){
        this.$message.success(fileid)
        this.tmpFiles.splice(index, 1)
        for (var i = 0; i < this.banners.length; i++) {
          if(this.banners[i].url == fileid) {
            this.banners.splice(i,1)
          }
        }
        for (var i = 0; i < this.storeImages.length; i++) {
          if(this.storeImages[i] == fileid) {
            this.storeImages.splice(i,1)
          }
        }

        for (var i = 0; i < this.productImages.length; i++) {
          if(this.productImages[i] == fileid) {
            this.productImages.splice(i,1)
          }
        }

        for (var i = 0; i < this.promoteImages.length; i++) {
          if(this.promoteImages[i] == fileid) {
            this.promoteImages.splice(i,1)
          }
        }
      },

      onExceed() {
        alert("超出数量")
      },
      submitUpload(index) {
        switch (index) {
          case 1:
            this.$refs.video.submit();
            break
          case 2:
            this.$refs.image.submit();
            break
          case 3:
            this.$refs.storeImage.submit();
            break
          case 4:
            this.$refs.productImage.submit();
            break
          case 5:
            this.$refs.promoteImage.submit();
        }
      },

      infohtmlset(result) {
        console.log(result)
        this.address = result.address
        this.longitude = result.point.lng
        this.latitude = result.point.lat
        this.mapVisible = false
      }
    },
    mounted() {
      this.center.lng = 116.240
      this.center.lat = 40.047
      this.zoom = 15

      if (this.id) {
        axios.get('http://localhost:8090/store?id='+this.id).then(resp => {

          // let tmpFiles = resp.data.split('---|---')[1]
          console.log(resp.data)
          // console.log(tmpFiles)
          let store = JSON.parse(resp.data.resp_data).data
          this.storeName = store.storeName
          this.storeDesc = store.storeDesc
          this.startTime = store.startTime
          this.endTime = store.endTime
          this.discount = store.discount
          this.discountDate = [store.discountStart, store.discountEnd]
          this.storeType = store.storeType
          // TODO 图片显示
          this.banners = store.banners,
          this.storeImages = store.storeImages,
          this.productImages = store.productImages,
          this.promoteImages = store.promoteImages,
          this.longitude = store.longitude
          this.latitude = store.latitude
          this.address = store.address
          this.merchantName = store.merchantName
          this.merchantPhone = store.merchantPhone
          this.merchantBankCard = store.merchantBankCard

          this.tmpFiles = resp.data.file_list
          for (var i = 0; i < this.tmpFiles.length; i ++) {
            if (this.tmpFiles[i].download_url.indexOf('mp4')!=-1){
              this.tmpFiles[i].isVideo = true
              // this.videoList.push({name:'a.mp4',url: this.tmpFiles[i].download_url})
            }else {
              this.tmpFiles[i].isVideo = false
            }
          }
        })
      }
    }
  }
</script>

<style>
  .el-row {
    margin-bottom: 20px;

    &:last-child {
      margin-bottom: 0;
    }
  }

  .el-col {
    border-radius: 4px;
  }

  .title {
    margin-top: 10px;
  }


  .map {
    width: 100%;
    height: 300px;
  }
</style>
