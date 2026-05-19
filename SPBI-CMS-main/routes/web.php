<?php

use Illuminate\Support\Facades\Route;
use App\Http\Controllers\UploadController;
use App\Http\Controllers\JabatanController;
use App\Http\Controllers\PenggunaController;
use App\Http\Controllers\AuthController;

/*
|--------------------------------------------------------------------------
| Web Routes
|--------------------------------------------------------------------------
|
| Here is where you can register web routes for your application. These
| routes are loaded by the RouteServiceProvider and all of them will
| be assigned to the "web" middleware group. Make something great!
|
*/

Route::get('/', [AuthController::class, 'loginView']);
Route::get('/login', [AuthController::class, 'loginView'])->name('login');
Route::get('/beranda', [AuthController::class, 'beranda'])->name('beranda');
Route::get('/jabatan', [JabatanController::class, 'index'])->name('jabatan');
Route::get('/jabatan/add', [JabatanController::class, 'add'])->name('jabatan.add');
Route::post('/jabatan/add', [JabatanController::class, 'simpan'])->name('jabatan.simpan');
Route::get('/jabatan/detail/{id}', [JabatanController::class, 'detail'])->name('jabatan.detail');
Route::get('/jabatan/data', [JabatanController::class, 'data'])->name('jabatan.data');
Route::post('/jabatan/delete', [JabatanController::class, 'delete'])->name('jabatan.delete');
Route::get('/pengguna', [PenggunaController::class, 'index'])->name('pengguna');
Route::get('/pengguna/data', [PenggunaController::class, 'data'])->name('pengguna.data');
Route::get('/pengguna/add', [PenggunaController::class, 'add'])->name('pengguna.add');
Route::post('/pengguna/add', [PenggunaController::class, 'simpan'])->name('pengguna.simpan');
Route::get('/pengguna/edit/{id}', [PenggunaController::class, 'edit'])->name('pengguna.edit');
Route::post('/pengguna/update/{id}', [PenggunaController::class, 'update'])->name('pengguna.update');
Route::post('/pengguna/delete', [PenggunaController::class, 'delete'])->name('pengguna.delete');
Route::get('/unggah-data',['as' => 'unggah-data', function () {
    return view('unggah');
}]);
Route::get('/unggah-data/neraca',['as' => 'unggah-data', function () {
    return view('unggah-neraca');
}]);
Route::get('/unggah-data/log',['as' => 'unggah-data', function () {
    return view('unggah-log');
}]);
Route::post('/upload', [UploadController::class, 'upload'])->name('posts.upload');
Route::get('/upload-history', [UploadController::class, 'uploadHistory'])->name('posts.upload-history');
Route::post('/upload-neraca', [UploadController::class, 'neraca'])->name('posts.upload-neraca');
// Route::post('/upload-neraca-history', [UploadController::class, 'neracaHistory'])->name('posts.upload-neraca-history');
Route::post('/login', [AuthController::class, 'login'])->name('posts.login');
Route::get('/logout', [AuthController::class, 'logout'])->name('get.logout');
Route::get('/forgot', [PenggunaController::class, 'forgot'])->name('forgot');
Route::post('/forgot', [PenggunaController::class, 'kirim'])->name('forgot.kirim');
Route::get('/reset', [PenggunaController::class, 'reset'])->name('reset');
Route::post('/reset-kirim', [PenggunaController::class, 'resetKirim'])->name('reset.kirim');