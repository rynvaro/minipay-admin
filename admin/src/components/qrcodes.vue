<template>
  <div>
    <el-button @click="downloadAll">一键下载</el-button>
    <div class="qrcode">
        <div class="qrcodeItem" v-for="(item, index) in stores" >
          <div :id="item._id" :ref="item._id"></div>
          <div>
            <span>{{item.storeName}}</span>
            <el-button type="text" @click="downloadImage(index)">下载</el-button>
          </div>

        </div>
    </div>
  </div>
</template>

<script>
  import QRCode from 'qrcodejs2'
  export default {
    name: 'qrcodes',
    data() {
      return {
        stores: this.$route.params.stores,
      }
    },
    methods: {
      downloadImage(index) {
        let myCanvas = document.getElementById(this.stores[index]._id).getElementsByTagName('canvas');
        let a = document.createElement('a')
        a.href = myCanvas[0].toDataURL('image/png')
        if (!this.stores[index].storeName) {
          a.download = this.stores[index]._id
        }else {
          a.download = this.stores[index].storeName
        }
        a.click()
      },
      downloadAll() {
        for (var i = 0; i<this.stores.length; i++ ) {
          let myCanvas = document.getElementById(this.stores[i]._id).getElementsByTagName('canvas');
          let a = document.createElement('a')
          a.href = myCanvas[0].toDataURL('image/png')
          if (!this.stores[i].storeName) {
            a.download = this.stores[i]._id
          }else {
            a.download = this.stores[i].storeName
          }
          a.click()
        }
      }
    },
    mounted() {
      for (var i = 0; i< this.stores.length; i++) {
        new QRCode(this.stores[i]._id, {
          width: 150,
          height: 150,
          text: this.stores[i]._id
        })
      }
    }
  }
</script>

<style>
  .qrcode {
    width: 100%;
    display: flex;
    justify-content: space-around;
    flex-wrap: wrap;
    align-items: center;
    border-radius: 10px;
  }

  .qrcodeItem {
    height: 200px;
    width: 200px;
    margin: 20px;
    background-color: #F6F6F6;
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
  }
</style>
