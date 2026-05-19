<!DOCTYPE html>
<html lang="{{ str_replace('_', '-', app()->getLocale()) }}">
    <head>
        <meta charset="utf-8">
        <meta name="viewport" content="width=device-width, initial-scale=1">

        <title>Info Pangan</title>


        <!-- Custom fonts for this template-->
        <link href="{{asset('vendor/fontawesome-free/css/all.min.css')}}" rel="stylesheet" type="text/css">
        <link href="https://fonts.googleapis.com/css2?family=Nunito:ital,wght@0,200..1000;1,200..1000&family=Plus+Jakarta+Sans:ital,wght@0,200..800;1,200..800&display=swap" rel="stylesheet">

        <!-- Custom styles for this template-->
        <link href="{{asset('css/sb-admin-2.css')}}" rel="stylesheet">
        <link href="{{asset('css/custom.css')}}" rel="stylesheet">
    </head>
    <body>

        <div class="container">

            <!-- Outer Row -->
            <div class="row justify-content-center">

                <div class="col-xl-10 col-lg-12 col-md-9">

                    <div style="background-color:unset" class="card o-hidden border-0 my-5">
                        <div class="card-body p-0">
                            <!-- Nested Row within Card Body -->
                            <div class="row">
                                <div class="col-lg-12">
                                    <div class="p-5">
                                        <div class="text-center">
                                            <div>
                                                <img width="80" src="{{asset('img/logo-sulsel.png')}}">
                                                <img width="80" src="{{asset('img/logo-bi.png')}}">
                                            </div>
                                            <div class="login-welcome text-gray-900">RESET PASSWORD</div>
                                            <div class="mb-4">Masukkan Password Baru
</div>
                                        </div>
                                        <div class="row">
                                            <div class="col-lg-3"></div>
                                            <div class="col-lg-6">
                                                @if (session('success'))
                                                    <div class="alert alert-success">{{ session('success') }}</div>
                                                @endif

                                                @if (session('error'))
                                                    <div class="alert alert-danger">{{ session('error') }}</div>
                                                @endif
                                                <div id="error_message" class="alert alert-danger" style="display:none">Password Do Not Match!</div>
                                                <form onsubmit="return validatePasswords(event)" class="user" action="{{route('reset.kirim')}}" method="post">
                                                @csrf
                                                <input type="hidden" name="code" value="{{ request()->query('code') }}">
                                                    <div class="form-group">
                                                        Password Baru
                                                        <div class="input-icons">
                                                            <i class="far fa-envelope icon"></i>
                                                            <input type="password" name="password" class="form-control form-login"
                                                                id="password" required aria-describedby="emailHelp">
                                                        </div>
                                                    </div>
                                                    <div class="form-group">
                                                        Retype Password Baru
                                                        <div class="input-icons">
                                                            <i class="far fa-envelope icon"></i>
                                                            <input type="password" name="password_re" class="form-control form-login"
                                                                id="confirm_password" required aria-describedby="emailHelp">
                                                        </div>
                                                    </div>
                                                    <button type="submit" href="{{route('reset.kirim')}}" class="btn btn-primary btn-custom-primary btn-block">
                                                        Kirim Link
</button>
                                                </form>
                                            </div>
                                            <div class="col-lg-3"></div>
                                        </div>
                                    </div>
                                </div>
                            </div>
                        </div>
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
    document.getElementById("confirm_password").addEventListener("input", function() {
        let password = document.getElementById("password").value;
        let confirmPassword = document.getElementById("confirm_password").value;
        let errorMessage = document.getElementById("error_message");

        if (password !== confirmPassword) {
                errorMessage.style.display = "block";
            } else {
                errorMessage.style.display = "none";
            }
    });
</script>
    </body>
</html>
