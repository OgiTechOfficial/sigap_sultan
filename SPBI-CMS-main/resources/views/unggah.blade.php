<!DOCTYPE html>
<html lang="en">

<head>

    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">
    <meta name="description" content="">
    <meta name="author" content="">

    <title>Info Pangan - Unggah Data</title>

    <!-- Custom fonts for this template-->
    <link href="{{asset('vendor/fontawesome-free/css/all.min.css')}}" rel="stylesheet" type="text/css">
    <link href="https://fonts.googleapis.com/css2?family=Nunito:ital,wght@0,200..1000;1,200..1000&family=Plus+Jakarta+Sans:ital,wght@0,200..800;1,200..800&display=swap" rel="stylesheet">

    <!-- Custom styles for this template-->
    <link href="{{asset('css/sb-admin-2.css')}}" rel="stylesheet">
    <link href="{{asset('css/custom.css')}}" rel="stylesheet">
    <!-- Custom styles for this page -->
    <link href="{{asset('vendor/datatables/dataTables.bootstrap4.min.css')}}" rel="stylesheet">

</head>

<body id="page-top">

    <!-- Page Wrapper -->
    <div id="wrapper">
    @include('menu')

        <!-- Content Wrapper -->
        <div id="content-wrapper" class="d-flex flex-column">

            <!-- Main Content -->
            <div id="content">

                <!-- Begin Page Content -->
                <div class="container-fluid">
                @include('topbar')
                    <div class="card" style="margin-bottom:30px">
                        <div class="row text-center">
                            <div class="col-lg-6" style="cursor: pointer;" onclick="window.location='{{url('unggah-data')}}';">
                                    <div style="padding:10px" class="border-bottom-blue">
                                        Unggah Data Harga
                                    </div>
                            </div>
                            <div class="col-lg-6" style="cursor: pointer;" onclick="window.location='{{url('unggah-data/neraca')}}';">
                                    <div style="padding:10px">
                                        Unggah Data Neraca
                                    </div>
                            </div>
                        </div>
                    </div>

                    <div class="row" style="margin-bottom: 30px;">
                        <div class="col-lg-6">
                            <div class="card" style="height:372px">
                                <div class="card-body">
                                    <div style="margin-bottom:24px" id="buttonContainer">
                                        <div class="weight60020" style="margin-bottom: 10px;">1. Download & Isi File CSV</div>
                                        <div class="custom50016">Upload daftar harga secara massal melalui template CSV yang disediakan.</div>
                                        <div><a class="btn btn-outline-light custom60014 downloadLink" style="margin-top: 24px;" href="{{asset('templates/price_national_template.csv')}}">Unduh Template CSV Nasional</a></div>
                                        <div><a class="btn btn-outline-light custom60014 downloadLink" style="margin-top: 24px;" href="{{asset('templates/price_province_template.csv')}}">Unduh Template CSV Provinsi</a></div>
                                        <div><a class="btn btn-outline-light custom60014 downloadLink" style="margin-top: 24px;" href="{{asset('templates/price_city_template.csv')}}">Unduh Template CSV Kab/Kota</a></div>
                                    </div>
                                </div>
                            </div>
                        </div>
                        <div class="col-lg-6">
                            <div class="card">
                                <div class="card-body">
                                    <div style="margin-bottom:24px">
                                        <div class="weight60020" style="margin-bottom: 10px;">2. Upload File Csv</div>
                                        <div class="custom50016">Pilih atau letakan file CSV .(csv) kamu disini.</div>
                                        Pilih Type Upload :
                                        <select name="upload_type" class="form-control" id="upload_type" onchange="getUrl(this)">
                                            <option value="price-city" selected>Kota</option>
                                            <option value="price-province">Provinsi</option>
                                            <option value="price-national">Nasional</option>
                                        </select>
                                        <div style="margin-top:24px">
                                            <form action="/file/post" class="dropzone dz-clickable border rounded bg-light p-3">
                                            @csrf
                                                <div class="dz-default dz-message text-center">
                                                    <div><span class="btn btn-primary btn-circle btn-lg"><i class="fas fa-file-import fa-flip-horizontal" style="font-size: 2rem;"></i></span></div>
                                                    <div class="btn btn-outline-light custom60014" style="margin-top:18px">Pilih atau Drag File</div>
                                                </div>
                                            </form>
                                        </div>
                                    </div>
                                </div>
                            </div></div>
                    </div>
                    <div class="row">
                        <div class="col-lg-12">
                            <div class="card">
                                <div class="card-body">
                                    <div class="row" style="margin-bottom:24px">
                                        <div class="col-lg-6 weight60020">3. Daftar File Upload</div>
                                    </div>
                                    <div class="table-responsive">
                                        <table class="table table-bordered" id="dataTable" width="100%" cellspacing="0">
                                            <thead>
                                                <tr>
                                                    <th>Waktu Upload</th>
                                                    <th>Nama File</th>
                                                    <th>Data Terunggah</th>
                                                </tr>
                                            </thead>
                                            <tbody>
                                                <tr>
                                                    <td>19 Januari 2024, 08:10 WIB</td>
                                                    <td>bulk01.xlsx</td>
                                                    <td>15/24</td>
                                                </tr>
                                            </tbody>
                                        </table>
                                    </div>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
                <!-- /.container-fluid -->

            </div>
            <!-- End of Main Content -->
        </div>
        <!-- End of Content Wrapper -->

    </div>
    <!-- End of Page Wrapper -->

    <!-- Scroll to Top Button-->
    <a class="scroll-to-top rounded" href="#page-top">
        <i class="fas fa-angle-up"></i>
    </a>

    <!-- Unduh Modal-->
    <div class="modal fade" id="unduhModal" tabindex="-1" role="dialog" aria-labelledby="unduhModalLabel"
        aria-hidden="true">
        <div class="modal-dialog" role="document">
            <div class="modal-content">
                <div class="modal-body">
                    <div class="row">
                        <div class="col-sm-2">
                            <span class="btn btn-success btn-circle btn-lg"><i class="fas fa-check"></i></span>
                        </div>
                        <div class="col-sm-8">
                            <div class="info-custom-hitam" style="margin:8px 0">Berhasil diunduh</div>
                            <div class="info-custom-grey">Template berhasil diunduh</div>
                        </div>
                        <div class="col-sm-2">
                            <button class="close" type="button" data-dismiss="modal" aria-label="Close">
                                <span aria-hidden="true">×</span>
                            </button>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </div>

    <!-- Success Modal-->
    <div class="modal fade" id="successModal" tabindex="-1" role="dialog" aria-labelledby="successModalLabel"
        aria-hidden="true">
        <div class="modal-dialog modal-dialog-centered" role="document">
            <div class="modal-content">
                <div class="modal-body text-center">
                    <div>
                        <span class="btn btn-success btn-circle btn-lg"><i class="fas fa-check-double"></i></span>
                    </div>
                    <div class="info-custom-hitam" style="margin:8px 0">Data Berhasil Diupload</div>
                    <button data-dismiss="modal" class="btn btn-primary btn-block">Tutup</button>
                </div>
            </div>
        </div>
    </div>

    <!-- Bootstrap core JavaScript-->
    <script src="{{asset('vendor/jquery/jquery.min.js')}}"></script>
    <script src="{{asset('vendor/bootstrap/js/bootstrap.bundle.min.js')}}"></script>

    <!-- Core plugin JavaScript-->
    <script src="{{asset('vendor/jquery-easing/jquery.easing.min.js')}}"></script>

    <!-- Custom scripts for all pages-->
    <script src="{{asset('js/sb-admin-2.min.js')}}"></script>

    <!-- Page level plugins -->
    <script src="{{asset('vendor/datatables/jquery.dataTables.min.js')}}"></script>
    <script src="{{asset('vendor/datatables/dataTables.bootstrap4.min.js')}}"></script>
    <link rel="stylesheet" href="https://unpkg.com/dropzone@5/dist/min/dropzone.min.css" />
    <script src="https://unpkg.com/dropzone@5/dist/min/dropzone.min.js"></script>
    <!-- Page level custom scripts -->
    <!-- <script src="{{asset('js/demo/datatables-demo.js')}}"></script> -->
    <script>
        Dropzone.autoDiscover = false;

        // Dropzone configuration

        var myDropzone = new Dropzone(".dropzone", {
            // url: urlUpload,
            paramName: "file",
            maxFilesize: 5, // MB
            maxFiles: 1,
            acceptedFiles: '.csv, application/vnd.openxmlformats-officedocument.spreadsheetml.sheet, application/vnd.ms-excel',
            dictDefaultMessage: "Pilih atau Drag File",
            clickable: true
        });
        myDropzone.on("processing", function(file) {
            var value = $('#upload_type').find(":selected").val();
            this.options.url = "<?php echo url('upload');?>?upload-type="+value;
        });
        myDropzone.on("success", function(file) {
            $("#successModal").modal('show');
        });
    </script>
    <script>
        document.getElementById('buttonContainer').addEventListener('click', function(event) {
            if (event.target.classList.contains('downloadLink')) {
                console.log('Button clicked:', event.target.textContent);
                $("#unduhModal").modal('show');
            }
        });
        $(document).ready(function() {
            $('#dataTable').DataTable({
        processing: true,
        serverSide: true,
        ajax: "<?php echo url('upload-history?module=PRICE');?>",
                columns: [
                    { data: 'createdAt' },
                    { data: 'fileName' },
                    { data: 'rowTotal' }
                ]
            });
        });
    </script>
</body>

</html>
