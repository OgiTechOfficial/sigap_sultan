<!-- Sidebar -->
<ul style="background-color:#fff" class="navbar-nav sidebar sidebar-dark accordion" id="accordionSidebar">

    <!-- Sidebar - Brand -->
    <div class="sidebar-top-custom">
        <div class="sidebar-brand-custom">
            <div>
                <img width="38" src="{{asset('img/logo-sulsel.png')}}">
                <img width="38" src="{{asset('img/logo-bi.png')}}">
            </div>
            <div><span class="color-white size-sultan">SIGAP</span><span class="color-yellow size-sultan">SULTAN</span>
            <br/><span class="color-white weight700">SISTEM INFORMASI</span>
            <br/><span class="color-white weight700">HARGA DAN PASOKAN PANGAN</span>
            <br/><span class="color-yellow weight700">SULAWESI SELATAN</span>
            </div>
        </div>
    </div>

    <!-- Divider -->
    <hr class="sidebar-divider my-0">

    <!-- Nav Item - Dashboard -->
    <li class="nav-custom nav-item-custom @if(Route::current()->getName() == 'beranda')
    active-custom
    @endif">
        <a class="@if(Route::current()->getName() == 'beranda')
    nav-link-custom-active
    @else
    nav-link-custom
    @endif" href="{{url('beranda')}}">
            <i class="fas fa-home"></i>
            <span>Beranda</span></a>
    </li>
    <li class="nav-custom nav-item-custom-bt @if(Route::current()->getName() == 'unggah-data')
    active-custom
    @endif">
        <a class="@if(Route::current()->getName() == 'unggah-data')
    nav-link-custom-active
    @else
    nav-link-custom
    @endif" href="{{url('unggah-data')}}">
            <i class="fas fa-file-import fa-flip-horizontal"></i>
            <span>Unggah Data</span></a>
    </li>
    <li class="nav-custom nav-item-custom-bt @if(Route::current()->getName() == 'jabatan')
    active-custom
    @endif">
        <a class="@if(Route::current()->getName() == 'jabatan')
    nav-link-custom-active
    @else
    nav-link-custom
    @endif" href="{{url('jabatan')}}">
            <i class="fas fa-book"></i>
            <span>Jabatan</span></a>
    </li>
    <li class="nav-custom nav-item-custom-bt @if(Route::current()->getName() == 'pengguna')
    active-custom
    @endif">
        <a class="@if(Route::current()->getName() == 'pengguna')
    nav-link-custom-active
    @else
    nav-link-custom
    @endif" href="{{url('pengguna')}}">
            <i class="fas fa-user"></i>
            <span>Pengguna</span></a>
    </li>
</ul>
<!-- End of Sidebar -->