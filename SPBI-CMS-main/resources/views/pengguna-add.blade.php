<!DOCTYPE html>
<html lang="en">

<head>

    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">
    <meta name="description" content="">
    <meta name="author" content="">

    <title>Info Pangan - Pengguna</title>

    <!-- Custom fonts for this template-->
    <link href="{{asset('vendor/fontawesome-free/css/all.min.css')}}" rel="stylesheet" type="text/css">
    <link href="https://fonts.googleapis.com/css2?family=Nunito:ital,wght@0,200..1000;1,200..1000&family=Plus+Jakarta+Sans:ital,wght@0,200..800;1,200..800&display=swap" rel="stylesheet">

    <!-- Custom styles for this template-->
    <link href="{{asset('css/sb-admin-2.css')}}" rel="stylesheet">
    <link href="{{asset('css/custom.css')}}" rel="stylesheet">

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
                    
                    <div class="row">
                        <div class="col-lg-12">
                            <div class="card">
                                <div class="card-body"> 
                                    <div class="row" style="margin-bottom:24px">
                                        <div class="col-lg-6 weight60020">Buat Pengguna Baru</div>
                                    </div>
                                    <div class="row">
                                        <div class="col-lg-12">
                                            <form class="user" id="submitPosition" method="post" action="{{url('pengguna/add')}}">
                                            @csrf
                                                <div class="form-group row">
                                                    <div class="col-sm-6">
                                                        Username
                                                        <div>
                                                            <input type="text" name="username" required class="form-control"
                                                                placeholder="Masukkan Username">
                                                        </div>
                                                    </div>
                                                    <div class="col-sm-6">
                                                        Nama
                                                        <div>
                                                            <input type="text" name="nama" required class="form-control"
                                                                placeholder="Masukkan Nama">
                                                        </div>
                                                    </div>
                                                </div>
                                                <div class="form-group row">
                                                    <div class="col-sm-6">
                                                        Email
                                                        <div>
                                                            <input type="email" name="email" required class="form-control"
                                                                placeholder="Masukkan Email">
                                                        </div>
                                                    </div>
                                                    <div class="col-sm-6">
                                                        Jabatan
                                                        <div>
                                                            <select name="jabatan" required class="form-control"
                                                                placeholder="Pilih Jabatan">
                                                                <?php foreach($jabatan as $jb):?>
                                                                <option value="{{ $jb->id.'|'.$jb->name }}">{{ $jb->name }}</option>
                                                                <?php endforeach;?>
                                                            </select>
                                                        </div>
                                                    </div>
                                                </div>
                                                <hr>
                                                <div class="float-right">
                                                    <a href="{{url('jabatan')}}" class="btn btn-outline-secondary btn-save-custom">
                                                        Batal
                                                    </a>
                                                    <button type="submit" class="btn btn-primary btn-custom-primary btn-save-custom">
                                                        Simpan
                                                    </button>
                                                </div>
                                            </form>
                                        </div>
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

    <!-- Success Modal-->
    <div class="modal fade" id="successModal" tabindex="-1" role="dialog" aria-labelledby="successModalLabel"
        aria-hidden="true">
        <div class="modal-dialog modal-dialog-centered" role="document">
            <div class="modal-content">
                <div class="modal-body text-center">
                    <div>
                        <span class="btn btn-success btn-circle btn-lg"><i class="fas fa-check-double"></i></span>
                    </div>
                    <div class="info-custom-hitam" style="margin:8px 0">Pengguna berhasil disimpan</div>
                    <div class="info-custom-grey" style="margin-bottom:24px">Pengguna berhasil disimpan</div>
                    <button data-dismiss="modal" class="btn btn-primary btn-block">Tutup</button>
                </div>
            </div>
        </div>
    </div>

    <!-- Gagal Modal-->
    <div class="modal fade" id="gagalModal" tabindex="-1" role="dialog" aria-labelledby="gagalModalLabel"
        aria-hidden="true">
        <div class="modal-dialog modal-dialog-centered" role="document">
            <div class="modal-content">
                <div class="modal-body text-center">
                    <div>
                        <span class="btn btn-danger btn-circle btn-lg"><i class="fas fa-times"></i></span>
                    </div>
                    <div class="info-custom-hitam" style="margin:8px 0">Pengguna gagal disimpan</div>
                    <div class="info-custom-grey" style="margin-bottom:24px">Pengguna tidak berhasil disimpan</div>
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
    <script>
        $(document).ready(function() {
            $('#submitPosition').submit(function(event) {
                event.preventDefault(); // Prevent the default form submission

                var formData = $(this).serialize(); // Serialize the form data, including the CSRF token

                $.ajax({
                url: '{{ route("pengguna.simpan") }}', // Laravel route for form submission
                method: 'POST',
                data: formData, // Send the form data
                success: function(response) {
                    $("#successModal").modal('show');
                },
                error: function(xhr, status, error) {
                    $("#gagalModal").modal('show');
                }
                });
            });
        });
        $('#successModal').on('hidden.bs.modal', function () {
            // Redirect to another page after the modal closes
            window.location.href = '<?php echo url('pengguna');?>'; // Replace with your desired URL
        });
    </script>
</body>

</html>