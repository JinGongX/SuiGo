<template>
  <el-row style="width: 86%;margin: 0 auto;">
       <div class="idx_menu_top">
        <router-link to="./Blogs">
          <div class="idx_menu_top_logo" >
            <svg t="1724676707084" class="icon" viewBox="0 0 1152 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="1475" width="32" height="32"><path d="M583.807877 778.756066c-84.099572 0-152.514558-70.908769-152.514558-158.092758s68.414986-158.059945 152.514558-158.059945 152.481745 70.908769 152.481744 158.092758-68.414986 158.059945-152.481744 158.059945z m0-261.453527c-53.944478 0-97.848195 46.364688-97.848196 103.360769s43.903717 103.393582 97.848196 103.393583 97.68413-46.364688 97.68413-103.36077-43.772466-103.393582-97.68413-103.393582z" fill="#ffffff" p-id="1476"></path><path d="M234.284411 1023.967187a109.16866 109.16866 0 0 1-109.16866-109.004596v-607.03944H179.814926v607.03944a54.403859 54.403859 0 0 0 54.338233 54.338234l685.462373-3.084417a54.436672 54.436672 0 0 0 54.469485-54.371046V311.073194h54.699176v600.772168a109.16866 109.16866 0 0 1-109.037409 109.16866z" fill="#ffffff" p-id="1477"></path><path d="M27.365994 419.060588a27.333181 27.333181 0 0 1-15.258018-50.006924L530.191528 18.873939a108.676466 108.676466 0 0 1 123.409478 0.689072l487.469077 337.251425a27.346306 27.346306 0 0 1-31.10667 44.986544L622.42871 64.516742a54.371046 54.371046 0 0 0-61.688332-0.328129L42.656826 414.368337a27.365994 27.365994 0 0 1-15.290832 4.692251z" fill="#ffffff" p-id="1478"></path></svg>          
            <span>首页</span>
          </div>
        </router-link> 
        <router-link to="./Blogs">
          <div class="idx_menu_top_logo" >
            <svg t="1724681892147" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="14205" width="32" height="32"><path d="M904 687.04a24 24 0 0 0-24 24v139.04a91.68 91.68 0 0 1-90.4 92.64H238.4a92.64 92.64 0 0 1-92.48-92.64V171.2a92.48 92.48 0 0 1 92.48-92.48h551.2A91.52 91.52 0 0 1 880 171.2v245.76a24 24 0 1 0 48 0V171.2a139.68 139.68 0 0 0-138.4-140.48H238.4a140.64 140.64 0 0 0-140.48 140.48v678.88a140.64 140.64 0 0 0 140.48 140.64h551.2A139.68 139.68 0 0 0 928 850.08v-139.04a24 24 0 0 0-24-24z" fill="#ffffff" p-id="14206"></path><path d="M706.08 288h-384a24 24 0 0 0 0 48h384a24 24 0 0 0 0-48zM705.12 480h-384a24 24 0 0 0 0 48h384a24 24 0 0 0 0-48zM1009.92 402.88a24.32 24.32 0 0 0-33.76 3.68l-361.44 448a23.84 23.84 0 0 0 3.52 33.6 23.84 23.84 0 0 0 33.76-3.52l361.6-448a24 24 0 0 0-3.68-33.76z" fill="#ffffff" p-id="14207"></path></svg>
              <span>随记</span>
          </div>
        </router-link>  
        <router-link to="./Chatsui">
          <div class="idx_menu_top_logo" >
             <svg t="1724681758292" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="11616" width="32" height="32"><path d="M511.702003 0C230.045029 0 0.303497 207.661769 0.303497 463.063026c0 115.168763 46.5455 224.967583 131.878583 310.299666V994.749314c0 11.337878 6.562929 21.481769 16.707821 26.255718 4.176955 1.789981 8.35391 2.983968 12.530865 2.983968 6.563929 0 13.127859-2.386974 17.901808-6.563929L345.213792 883.160513c12.531865-10.143891 14.321846-28.045699 4.176955-40.577564-10.143891-12.530865-28.045699-14.321846-40.577564-4.176955L189.467465 934.479962V761.427821c0-7.757917-2.983968-15.514833-8.950904-20.885776-78.769154-75.784186-122.329686-174.245128-122.329686-277.479019 0-223.177603 203.484814-405.180647 453.515128-405.180648 250.029314 0 453.515128 182.003045 453.515128 405.180648s-203.485814 405.179647-453.515128 405.179647h-3.579961c-16.112827 0-28.643692 12.530865-28.643693 28.642692 0 16.111827 12.530865 29.239686 28.642693 29.239686h4.176955c282.253968 0 511.398506-207.661769 511.398506-463.062025C1023.696503 207.661769 793.954971 0 511.702003 0zM231.238016 415.920532c-12.530865 12.531865-19.691788 29.836679-19.691788 47.739487 0 17.304814 7.160923 35.206622 19.691788 47.737487 12.531865 12.531865 29.836679 19.692788 47.739487 19.692789 17.304814 0 35.206622-7.160923 47.737487-19.692789 12.531865-12.530865 19.692788-29.836679 19.692789-47.738487 0-17.304814-7.160923-35.206622-19.692789-47.738487-12.530865-12.530865-29.836679-19.691788-47.738487-19.691788-17.901808 0-35.206622 7.159923-47.738487 19.691788z m328.202474 95.476974c12.530865-12.530865 19.691788-29.836679 19.691789-47.738487 0-17.304814-7.160923-35.206622-19.691789-47.738487-12.531865-12.530865-29.836679-19.691788-47.739487-19.691788-17.304814 0-35.206622 7.159923-47.737487 19.691788-12.531865 12.531865-19.692788 29.836679-19.692788 47.739487 0 17.304814 7.160923 35.206622 19.692788 47.737487 12.530865 12.531865 29.836679 19.692788 47.738487 19.692789 17.901808 0 35.206622-7.757917 47.738487-19.692789z m233.321494 0c12.530865-12.530865 19.691788-29.836679 19.691788-47.738487 0-17.304814-7.160923-35.206622-19.691788-47.738487-12.531865-12.530865-29.836679-19.691788-47.739487-19.691788-17.304814 0-35.206622 7.159923-47.737487 19.691788-12.531865 12.531865-19.692788 29.836679-19.692789 47.739487 0 17.304814 7.160923 35.206622 19.692789 47.737487 12.530865 12.531865 29.836679 19.692788 47.738487 19.692789 17.901808 0 35.206622-7.757917 47.738487-19.692789z" fill="#ffffff" p-id="11617"></path></svg>
              <span>随问</span>
          </div>
        </router-link> 
        <router-link to="./ExcalidrawWrapper" target="_blank" rel="opener"> 
          <div class="idx_menu_top_logo">
            <svg t="1724680891046" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="2719" width="32" height="32"><path d="M314.14272 526.12096h56.53504V695.6032h-56.53504V526.12096z m169.59488-169.472h56.5248V695.6032h-56.5248V356.64896z m169.59488 112.97792h56.53504V695.6032h-56.53504V469.62688z" fill="#ffffff" p-id="2720"></path><path d="M935.99744 436.56192c0-332.78976-45.21984-320.07168-452.2496-320.07168-407.02976 0-395.7248-12.70784-395.7248 320.07168 0 306.25792-8.68352 351.744 307.03616 357.07904L259.71712 931.86048c-7.168 7.424-7.168 19.2 0 26.624l0.37888 0.36864a18.2272 18.2272 0 0 0 25.74336-0.36864L446.6688 794.18368c12.32896 0.07168 23.88992 0.19456 37.08928 0.19456 3.47136 0 6.42048-0.0512 9.84064-0.0512v151.11168a18.42176 18.42176 0 0 0 36.84352 0V794.14272c15.72864-0.11264 31.44704-0.29696 47.16544-0.59392l160.93184 164.87424a18.25792 18.25792 0 0 0 25.81504 0.47104l0.4608-0.47104a19.4048 19.4048 0 0 0 0-26.94144l-136.00768-139.264c272.8448-10.28096 307.18976-67.6864 307.18976-355.65568zM569.088 737.9456a15.91296 15.91296 0 0 0-1.97632 0c-17.55136-0.03072-35.51232-0.07168-55.0912-0.07168-9.984 0-18.6368 0.02048-28.08832 0.04096-0.06144 0-0.11264-0.04096-0.18432-0.04096-0.06144 0 0 0.04096-0.06144 0.04096-342.19008 0.47104-339.12832 5.29408-339.12832-282.49088 0-295.7824-32.62464-282.44992 339.18976-282.44992s395.7248-13.39392 395.7248 282.44992c0 279.71584 3.42016 282.99264-310.38464 282.5216zM456.38656 59.99616h111.24736a18.5344 18.5344 0 1 1 0 37.0688H456.38656a18.54464 18.54464 0 0 1 0-37.0688z" fill="#ffffff" p-id="2721"></path></svg>     
                   <span>白板</span>
          </div>
        </router-link>
        <router-link to="./BookInfo"> 
          <div class="idx_menu_top_logo">
            <svg t="1724681226257" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="7121" width="32" height="32"><path d="M897.6384 448.512H368.64a140.4928 140.4928 0 1 1 0-281.1904h528.9984a20.48 20.48 0 1 1 0 40.96H368.64a99.5328 99.5328 0 1 0 0 199.0656h528.9984a20.48 20.48 0 0 1 0 40.96z" fill="#ffffff" p-id="7122"></path><path d="M835.584 444.416a20.48 20.48 0 0 1-16.7936-8.8064A222.8224 222.8224 0 0 1 819.2 177.152a20.48 20.48 0 0 1 32.3584 25.1904 178.7904 178.7904 0 0 0 0 209.92 20.48 20.48 0 0 1-16.7936 32.1536zM897.6384 940.032H368.64a140.4928 140.4928 0 1 1 0-280.9856h528.9984a20.48 20.48 0 0 1 0 40.96H368.64a99.5328 99.5328 0 1 0 0 199.0656h528.9984a20.48 20.48 0 0 1 0 40.96z" fill="#ffffff" p-id="7123"></path><path d="M835.584 935.936a20.48 20.48 0 0 1-16.7936-8.8064 223.0272 223.0272 0 0 1 0-258.4576 20.48 20.48 0 0 1 32.3584 25.1904 178.7904 178.7904 0 0 0 0 209.92 20.48 20.48 0 0 1-16.7936 32.1536zM655.36 694.272H126.3616a20.48 20.48 0 0 1 0-40.96H655.36a99.5328 99.5328 0 0 0 0-199.0656H126.3616a20.48 20.48 0 1 1 0-40.96H655.36a140.4928 140.4928 0 0 1 0 280.9856z" fill="#ffffff" p-id="7124"></path><path d="M188.416 690.176a20.48 20.48 0 0 1-16.7936-32.1536 178.7904 178.7904 0 0 0 0-209.92A20.48 20.48 0 0 1 204.8 422.912a222.8224 222.8224 0 0 1 0 258.4576 20.48 20.48 0 0 1-16.384 8.8064z" fill="#ffffff" p-id="7125"></path></svg>
                        <span>读书</span>
          </div>
        </router-link>
        <router-link to="./SoftInfo"> 
          <div class="idx_menu_top_logo">
             <svg t="1724681383017" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="10461" width="32" height="32"><path d="M614.4 899.2768H98.0992A98.304 98.304 0 0 1 0 801.1776V204.8a98.304 98.304 0 0 1 98.0992-97.6896h743.8336A98.304 98.304 0 0 1 940.032 204.8v215.04a20.48 20.48 0 1 1-40.96 0V204.8a57.1392 57.1392 0 0 0-57.1392-57.1392H98.0992A57.1392 57.1392 0 0 0 40.96 204.8v595.968a57.1392 57.1392 0 0 0 57.1392 57.1392H614.4a20.48 20.48 0 1 1 0 40.96z" fill="#ffffff" p-id="10462"></path><path d="M919.552 350.6176H20.48a20.48 20.48 0 0 1 0-40.96h899.072a20.48 20.48 0 0 1 0 40.96z" fill="#ffffff" p-id="10463"></path><path d="M678.2976 230.4m-19.8656 0a19.8656 19.8656 0 1 0 39.7312 0 19.8656 19.8656 0 1 0-39.7312 0Z" fill="#ffffff" p-id="10464"></path><path d="M744.8576 230.4m-19.8656 0a19.8656 19.8656 0 1 0 39.7312 0 19.8656 19.8656 0 1 0-39.7312 0Z" fill="#ffffff" p-id="10465"></path><path d="M809.984 230.4m-19.8656 0a19.8656 19.8656 0 1 0 39.7312 0 19.8656 19.8656 0 1 0-39.7312 0Z" fill="#ffffff" p-id="10466"></path><path d="M184.32 717.824a20.48 20.48 0 0 1-15.1552-6.7584l-85.6064-94.4128a20.48 20.48 0 0 1 0-27.8528l86.4256-91.7504a20.48 20.48 0 0 1 29.0816 0 20.48 20.48 0 0 1 0 28.8768l-72.4992 77.2096 73.728 80.2816a20.48 20.48 0 0 1-1.2288 28.8768 20.48 20.48 0 0 1-14.7456 5.5296zM395.0592 717.824a20.48 20.48 0 0 1-13.7216-5.5296 20.48 20.48 0 0 1-1.4336-28.8768l73.728-80.2816-73.5232-78.0288A20.48 20.48 0 0 1 409.6 497.0496l86.6304 91.7504a20.48 20.48 0 0 1 0 27.8528L409.6 711.0656a20.48 20.48 0 0 1-14.5408 6.7584zM253.3376 757.76a17.2032 17.2032 0 0 1-5.5296 0 20.48 20.48 0 0 1-14.1312-25.3952l75.9808-267.4688a20.48 20.48 0 0 1 39.5264 11.264l-76.1856 266.0352a20.48 20.48 0 0 1-19.6608 15.5648zM767.1808 772.096a110.7968 110.7968 0 1 1 110.592-110.7968 111.0016 111.0016 0 0 1-110.592 110.7968z m0-180.4288a69.8368 69.8368 0 1 0 69.632 69.632 69.8368 69.8368 0 0 0-69.632-69.632z" fill="#ffffff" p-id="10467"></path><path d="M778.24 916.8896a260.9152 260.9152 0 0 1-61.44-7.3728 20.48 20.48 0 0 1-15.5648-20.48 61.44 61.44 0 0 0-61.44-61.44 60.0064 60.0064 0 0 0-27.8528 7.168 20.48 20.48 0 0 1-24.7808-4.5056A253.5424 253.5424 0 0 1 532.48 733.184a20.48 20.48 0 0 1 9.216-23.3472 61.44 61.44 0 0 0 22.1184-83.968 59.1872 59.1872 0 0 0-20.48-20.48 20.48 20.48 0 0 1-8.3968-23.7568A257.024 257.024 0 0 1 593.92 485.1712a20.48 20.48 0 0 1 24.9856-3.8912 61.44 61.44 0 0 0 83.5584-22.7328 60.2112 60.2112 0 0 0 7.3728-28.4672 20.48 20.48 0 0 1 16.384-19.2512 251.0848 251.0848 0 0 1 112.4352 2.048 20.48 20.48 0 0 1 15.5648 20.48 61.44 61.44 0 0 0 89.7024 53.8624 20.48 20.48 0 0 1 24.576 4.5056A253.5424 253.5424 0 0 1 1024 589.824a20.48 20.48 0 0 1-9.216 23.3472A61.44 61.44 0 0 0 991.8464 696.32a61.44 61.44 0 0 0 20.48 20.48 20.48 20.48 0 0 1 8.3968 23.7568A256.2048 256.2048 0 0 1 962.56 838.0416a20.48 20.48 0 0 1-24.7808 3.6864 61.44 61.44 0 0 0-83.5584 22.7328 58.1632 58.1632 0 0 0-7.9872 28.0576 20.48 20.48 0 0 1-16.1792 19.2512 261.9392 261.9392 0 0 1-51.8144 5.12z m-36.864-44.2368a211.5584 211.5584 0 0 0 66.1504 1.4336 100.352 100.352 0 0 1 11.6736-30.1056A102.4 102.4 0 0 1 942.08 798.72a213.8112 213.8112 0 0 0 34.2016-56.7296 96.8704 96.8704 0 0 1-20.48-24.576 102.4 102.4 0 0 1 22.9376-129.8432 210.5344 210.5344 0 0 0-32.1536-57.7536A102.4 102.4 0 0 1 814.8992 450.56a199.4752 199.4752 0 0 0-66.1504-1.2288 102.4 102.4 0 0 1-11.4688 29.696 102.4 102.4 0 0 1-122.88 45.056 219.3408 219.3408 0 0 0-34.2016 56.7296 96.8704 96.8704 0 0 1 20.48 24.576 102.4 102.4 0 0 1-22.9376 129.8432 210.5344 210.5344 0 0 0 32.1536 57.7536 102.4 102.4 0 0 1 132.3008 79.6672z" fill="#ffffff" p-id="10468"></path></svg>
                       <span>软件</span>
          </div> 
        </router-link>
        <router-link to="./myBlogs"> 
          <div class="idx_menu_top_logo">
            <svg t="1724681979895" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="15395" width="32" height="32"><path d="M789.8 937.8H237.2c-76.8 0-139.4-62.6-139.4-139.4v-64.6c0-76.8 62.6-139.4 139.4-139.4H790c76.8 0 139.4 62.6 139.4 139.4v64.6c-0.2 76.8-62.6 139.4-139.6 139.4zM237.2 646.4c-48.2 0-87.4 39.2-87.4 87.4v64.6c0 48.2 39.2 87.4 87.4 87.4H790c48.2 0 87.4-39.2 87.4-87.4v-64.6c0-48.2-39.2-87.4-87.4-87.4H237.2z" p-id="15396" fill="#ffffff"></path><path d="M513.4 646.8c-153.6 0-278.4-125-278.4-278.4C235 215 360 90 513.4 90 667 90 792 215 792 368.4c0 153.6-125 278.4-278.6 278.4z m0-504.8c-124.8 0-226.4 101.6-226.4 226.4s101.6 226.4 226.4 226.4S740 493.4 740 368.4 638.4 142 513.4 142z" p-id="15397" fill="#ffffff"></path><path d="M337 394.4c-14.4 0-26-11.6-26-26 0-109.6 89.2-199 199-199 14.4 0 26 11.6 26 26s-11.6 26-26 26c-81 0-147 66-147 147 0 14.4-11.6 26-26 26z" fill="#ffffff" p-id="15398"></path></svg>
              <span>我的</span>
          </div>
        </router-link>
          <div class="idx_menu_top_logo" title="敬请期待">
            <svg t="1724682079454" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="16953" width="32" height="32"><path d="M298.666667 117.333333a181.333333 181.333333 0 0 1 181.333333 181.333334v181.333333H298.666667a181.333333 181.333333 0 1 1 0-362.666667z m117.333333 298.666667V298.666667a117.333333 117.333333 0 1 0-117.333333 117.333333h117.333333zM725.333333 117.333333a181.333333 181.333333 0 1 1 0 362.666667h-181.333333V298.666667A181.333333 181.333333 0 0 1 725.333333 117.333333z m-117.333333 298.666667H725.333333a117.333333 117.333333 0 1 0-117.333333-117.333333v117.333333zM298.666667 906.666667a181.333333 181.333333 0 1 1 0-362.666667h181.333333V725.333333a181.333333 181.333333 0 0 1-181.333333 181.333334z m117.333333-298.666667H298.666667a117.333333 117.333333 0 1 0 117.333333 117.333333v-117.333333z m309.333333 298.666667a181.333333 181.333333 0 0 1-181.333333-181.333334v-181.333333H725.333333a181.333333 181.333333 0 1 1 0 362.666667z m-117.333333-298.666667V725.333333a117.333333 117.333333 0 1 0 117.333333-117.333333h-117.333333z" fill="#ffffff" p-id="16954"></path></svg>
              <span>更多</span>
          </div>
        </div>
    </el-row>
  <!-- <el-row style="width: 86%;margin: 6px auto;text-align: center;">
    <el-col :span="24"><el-button-group class="ml-4">
       <el-button type="success" :icon="Files" @click="GetAllDatas(1)">全部随记</el-button> 
       <router-link to="./myBlogs"><el-button type="success" >我的随记</el-button></router-link>
     </el-button-group>
    </el-col>
  </el-row> -->
  <el-row style="width: 86%;margin: 0 auto;">
    <el-skeleton style="width: 240px" :loading="loading" animated :count="3" >
    <template #template> 
      <el-skeleton-item  variant="image" style="width: 240px; height: 240px" />
      <div style="padding: 14px">
        <el-skeleton-item variant="p" style="width: 50%" />
        <div
          style="
            display: flex;
            align-items: center;
            justify-items: space-between;
          "
        >
          <el-skeleton-item variant="text" style="margin-right: 16px" />
          <el-skeleton-item variant="text" style="width: 30%" />
        </div>
      </div> 
    </template>
  </el-skeleton>  
  <el-space wrap   style="margin-bottom: 0;width: 100%;display: grid;grid-template-columns: repeat(auto-fill, minmax(260px, 1fr));gap: 30px;" >
 
      <el-card  class="box-card" v-for="(o, index) in  blogData "   :key="o" :offset="index > 0 ? index%4 == 0? 0:1 :  0"   :body-style="{ padding: '0px' }"> 
        <img v-if="blogData[index].imgcover!=''"
          :src=blogData[index].imgcover
          class="image" 
        />
        <!-- <div  v-if="blogData[index].imgcover!=''" style="width: 100%; height: 144px;background-repeat: no-repeat;background-size:cover;background-position: center center;" :style="{backgroundImage:`url(`+blogData[index].imgcover+`)`}" ></div> -->
        <el-skeleton-item v-else variant="image" style="width: 100%; height: 144px" />
        <!-- <img v-else class="image"  
           src="../../assets/sysimg/Emay.png"
        /> -->
        <!-- "https://shadow.elemecdn.com/app/element/hamburger.9cf7b091-55e9-11e9-a976-7f4d0b07eef6.png" -->
        <div style="padding: 14px">
          <h3 class="blog-title" :title=blogData[index].title > <!--{{ blogData[index].account.name }} -->
            <router-link :to="{path:'/Testing',query:{id:blogData[index].id}}" >{{ blogData[index].title }}</router-link>
          </h3>
          <span class="blog-overview">
            <!-- {{ message.substring(0,51)+"..." }} -->
            {{ blogData[index].desc}}
          </span>
         
          <div class="bottom" style="margin-top: 6px;">
            <img alt="photo" class="photo" :title=blogData[index].account.name :src=blogData[index].account.photo /> 
           
            <time class="time"> {{ this.moments(blogData[index].created_on*1000).format("YYYY-MM-DD") }}</time>
            <span style="font-size: 13px;color:#999;">阅读 {{ blogData[index].views }} </span>
            <div class="crard_tab"><el-tag effect="plain" type="warning">{{ blogData[index].tag.name }}</el-tag></div>
            
            <!-- HH:mm <router-link :to="{path:'/Testing',query:{id:blogData[index].id}}"><el-button text class="button"><el-icon><View /></el-icon>查看详细</el-button></router-link> -->
          </div>
          
          <!-- <div class="foot_btn">
            <div class="crard_tab"><el-tag effect="plain" type="warning">{{ blogData[index].tag.name }}</el-tag></div>
            <router-link :to="{path:'/Testing',query:{id:blogData[index].id}}"  ><el-button type="success"  icon="View" circle title="查看详细" /></router-link> </div>
           -->
        </div>
      </el-card>
    <!-- </el-col> -->
  </el-space>
    <el-col align="middle" style="margin-top: 10px" >
      <el-button @click="GetDatas()" v-if="btdisabled==false" >加载数据..</el-button>
      <el-button  v-else  disabled>-显示完毕-</el-button> 
     
    </el-col>
  </el-row>
</template>
<script>
import { ref} from 'vue' 
import {
  View,Files,
} from '@element-plus/icons-vue'
import axios from 'axios'
import VueAxios from 'vue-axios'

import moment from "moment"
import { useRouter } from 'vue-router'


moment.locale('zh-cn')

export default {
  data() {
    return {
      form: {
        desc: ''
      },
      blogData:[],
      total:0,
      currentDate : ref(new Date()),
      message: "",
      page:1,
      btvalue:'加载数据..',
      btdisabled:false,
      moments:moment,
      cjtypes:0,
      loading: true,
      size:30
    }
  },
  mounted() {
    axios({
         method: "get",
         url: this.$globalHost+"api/v1/articles?state=1", //"http://127.0.0.1:2023/api/v1/articles?state=1",
         headers:{
            "token":sessionStorage.token
         }
        })
        .then((res) => {
          this.loading = false
         //test=res.data.data.content 
         console.log(res.data)
      this.blogData=res.data.data.lists
      this.total=res.data.data.total
         console.log("userid:"+sessionStorage.id)
         //datas.blogData=JSON.parse(JSON.stringify(res.data.data))  //eval(res.data.data).lists[7]
         if(this.blogData.length==this.total){
            this.btvalue="-显示完毕-"
            this.btdisabled=true
           }
          //  let tds=this.moments().unix(1690728099*1000)
          //  console.log("test:"+ this.moments(1691674049*1000).format("YYYY-MM-DD HH:mm:ss"))
         
    }) 
  },
  methods: {
     GetDatas(){
      var types=""
      if(this.cjtypes=="2"){
        types="&created_by="+sessionStorage.id
      }
       this.page=this.page+1
        //alert(this.page)
        axios({
           method: "get",
           url: this.$globalHost+"api/v1/articles?state=1&page="+this.page+types, //"http://127.0.0.1:2023/api/v1/articles?state=1&page="+this.page+types,
           headers:{
              "token":sessionStorage.token
           }
          })
          .then((res) => {
            this.loading = false
           //test=res.data.data.content 
           console.log(res.data)
           this.blogData.push(...res.data.data.lists)
           this.total=res.data.data.total
           if(this.blogData.length==this.total){
            this.btvalue="-显示完毕-"
            this.btdisabled=true
           }
           console.log(this.total) 
           
           }) 
    },
    GetAllDatas(sjtype){
      this.cjtypes=sjtype
      var types=""
      if(sjtype=="2"){
        types="&state=1?created_by="+sessionStorage.id
      }
      this.page=1
      axios({
         method: "get",
         url:  this.$globalHost+"api/v1/articles?state=1&"+types, //"http://127.0.0.1:2023/api/v1/articles?state=1&"+types,
         headers:{
            "token":sessionStorage.token
         }
        })
        .then((res) => { 
         this.blogData=res.data.data.lists
         this.total=res.data.data.total
         console.log("userid:"+sessionStorage.id) 
         if(this.blogData.length==this.total){
            this.btvalue="-显示完毕-"
            this.btdisabled=true
         }else{
           this.btvalue="加载数据.."
           this.btdisabled=false
         }
    }) 
    },
    onmenu(menuid){
      const router = useRouter();
      if(menuid==0){
        router.push("Blogs") 
      }else if(menuid==1){
        router.push("Blogs") 
      }else if(menuid==2){
        router.push("Chatsui") 
      }else if(menuid==4){
        router.push("ExcalidrawWrapper") 
      }
      
      
 
    }
  }
}
 
</script>
<!-- <script lang="ts" setup>
import { ref,reactive,onMounted, toRef } from 'vue' 
import {
  View,
} from '@element-plus/icons-vue'
import axios from 'axios'

const currentDate = ref(new Date())
var message= "你是我生命中最美的风景，每天都因为你的存在而变得更加灿烂。"
var datas=reactive({
  blogData:[]
})
onMounted(async()=>{
   await axios({
         method: "get",
         url: "http://127.0.0.1:2023/api/v1/articles",
         headers:{
            "token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6InRlc3QiLCJwYXNzd29yZCI6InRlc3QxMjM0NTYiLCJleHAiOjE2OTA5ODQ4MTYsImlzcyI6Imdpbi1ibG9nIn0.A7WGJ2G83kHbpRETAurEggm0gmXhinLU-zWEE-hizHk"
         }
        })
        .then((res) => {
         //test=res.data.data.content 
         console.log(res.data)
         
         datas.blogData=JSON.parse(JSON.stringify(res.data.data))  //eval(res.data.data).lists[7]
         
    }) 
    console.log("blogdata:"+datas.blogData)
    return {
      ...toRef(datas.blogData)
    }
 })
 //return {
 // ...toRef(datas.blogData)
 //}
      
</script> -->
<style>
.time {
  font-size: 14px;
  color: #999;
}

.bottom {
  margin-top: 13px;
  line-height: 12px;
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.button {
  padding: 0;
  min-height: auto;
}

.image {
  width: 100%;
  display: block;
  max-height:144px;
  object-fit: scale-down;
  min-height: 144px;
}
.blog-overview{
  color:#202124;
  font-size:14px;
  line-height: 22px;
  word-break: break-all;
  display: -webkit-box;
  overflow: hidden;
  text-overflow: ellipsis;
  -webkit-box-orient: vertical;
  -webkit-line-clamp: 3;
  margin-top:12px;
  min-height:66px;
}
.foot_btn{
    text-align:right;
  }

  .el-skeleton__image svg {
    color: var(--el-svg-monochrome-grey);
    fill: currentColor;
    width: 52%;
    height: 52%;
}
.blog-title{
  overflow:hidden;
  text-overflow: ellipsis;
  display: -webkit-box;
  -webkit-line-clamp: 1;
  -webkit-box-orient: vertical;
}
.crard_tab{
  float: left;
  margin-top: 4px;
}
.idx_menu_top{
  display: flex;
  /* display: grid; */
  flex-direction: row;
  column-gap: 50px;
  row-gap: 20px;
  justify-content: center;
  width:100%;
  margin-bottom: 30px;
  /* grid-template-columns: repeat(auto-fill, minmax(60px, 1fr)); */
}
.idx_menu_top_logo{ 
  width: 50px;
    height: 50px; 
    display: flex;
    flex-wrap: wrap;
    justify-content: center;
    border-radius: 50%;
    background-color: #6bdd88; 
    cursor: pointer;
}
.idx_menu_top_logo svg{
  width: 30px;
  height: 30px;
  margin: 10px;
}
.idx_menu_top_logo span{
  margin-top: -4px;
  display: block;
  text-align: center;
  font-size: 13px;
  color:#000;
}
</style>