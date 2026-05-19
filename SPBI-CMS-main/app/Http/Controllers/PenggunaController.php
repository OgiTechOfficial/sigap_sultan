<?php

namespace App\Http\Controllers;

use CURLFile;
use CURLStringFile;
use Illuminate\Http\Request;
use Illuminate\Support\Facades\Http;
use Illuminate\Support\Facades\DB;
use Yajra\DataTables\DataTables;
// use GuzzleHttp\Client;

class PenggunaController extends Controller
{
    public function index(Request $request)
    {
        return view('pengguna',[
            'active_menu' => 'pengguna'
        ]);
    }

    public function detail($id){
        $result = Http::get(env('API_URL').'/position/'.$id);
        $data = $result->object();
        
        $resultm = Http::get(env('API_URL').'/menu');
        // echo $result->ok();
        // die();
        if ($data->message != "Success"){
            abort(404);
        }else{
            return view('jabatan-detail',[
                'active_menu' => 'jabatan',
                'data' => $data->data,
                'datamenu' => $resultm->object()->data
            ]);
        }
    }

    public function data(Request $request){
        if ($request->ajax()) {
            $users = DB::table('prod.tm_user AS a')
            ->join("prod.tm_profile AS b", "b.user_id", "=", "a.id")
            ->select(DB::raw("a.id,username, b.nama_depan AS full_name,email, b.jabatan"));
            return DataTables::of($users)
                ->addColumn('action', function ($user) {
                    $url = url('pengguna/edit/'.$user->id);
                    return '<a href="'.$url.'" class="btn btn-sm btn-primary">Ubah</a>
                <button class="btn btn-sm btn-outline-secondary del-pos" data-id="'.$user->id.'">Hapus</button>';
                })
                ->rawColumns(['action']) // For rendering HTML
                ->make(true);
        }
    }
    
    public function add(Request $request)
    {
        $result = Http::get(env('API_URL').'/menu');
        $jabatan = Http::get(env('API_URL').'/position');
        return view('pengguna-add',[
            'active_menu' => 'pengguna',
            'menudata' => $result->object()->data,
            'jabatan' => $jabatan->object()->data
        ]);
    }

    public function simpan(Request $request)
    {
        $jabatan = explode("|", $request->jabatan);
        try {
            DB::beginTransaction();
        
            // Insert into the first table
            $userId = DB::table('prod.tm_user')->insertGetId([
                'client_id' => 1,
                'username' => $request->username,
                'email' => $request->email,
                'password' => hash('sha256', $request->email),
            ]);
        
            // Insert into the second table
            DB::table('prod.tm_profile')->insert([
                'client_id' => 1,
                'user_id' => $userId,
                'nama_depan' => $request->nama,
                'jabatan' => $jabatan[1],
                'position_id' => $jabatan[0],
            ]);
        
            // Commit the transaction
            DB::commit();
        
            return response()->json(['message' => 'Pengguna Berhasil Disimpan']);
        } catch (\Exception $e) {
            // Rollback the transaction on error
            DB::rollBack();
        
            return response()->json(['error' => 'Something went wrong', 'message' => 'Pengguna Gagal Disimpan'], 400);
        }
    }
    
    public function edit($id)
    {
        $result = Http::get(env('API_URL').'/menu');
        $jabatan = Http::get(env('API_URL').'/position');
        $detail = DB::table('prod.tm_user AS a')
        ->join("prod.tm_profile AS b", "b.user_id", "=", "a.id")
        ->select(DB::raw("a.id,username, b.nama_depan, b.nama_belakang,email, b.jabatan, b.position_id"))
        ->where("a.id", $id)
        ->first();
        return view('pengguna-edit',[
            'active_menu' => 'pengguna',
            'menudata' => $result->object()->data,
            'jabatan' => $jabatan->object()->data,
            'detail' => $detail
        ]);
    }

    public function update($id, Request $request)
    {
        $jabatan = explode("|", $request->jabatan);
        try {
            DB::beginTransaction();
        
            DB::table('prod.tm_user')
                ->where('id', $id)
                ->update([
                    'username' => $request->username,
                    'email' => $request->email
                ]);
        
            // Insert into the second table
            DB::table('prod.tm_profile')
            ->where('user_id', $id)
            ->update([
                'nama_depan' => $request->nama,
                'jabatan' => $jabatan[1],
                'position_id' => $jabatan[0],
            ]);
        
            // Commit the transaction
            DB::commit();
            
            session(['name' => $request->nama]);
        
            return response()->json(['message' => 'Pengguna Berhasil Disimpan']);
        } catch (\Exception $e) {
            // Rollback the transaction on error
            DB::rollBack();
        
            $queries = DB::getQueryLog();
            dd($queries);
            return response()->json(['error' => 'Something went wrong', 'message' => 'Pengguna Gagal Disimpan'], 400);
        }
    }

    public function delete(Request $request)
    {
        try {
            DB::beginTransaction();
        
            DB::table('prod.tm_user')
                ->where('id', $request->id)
                ->delete();
        
            // Insert into the second table
            DB::table('prod.tm_profile')
            ->where('user_id', $request->id)
            ->delete();
        
            // Commit the transaction
            DB::commit();
        
            return response()->json(['message' => 'Pengguna Berhasil Dihapus']);
        } catch (\Exception $e) {
            // Rollback the transaction on error
            DB::rollBack();
        
            $queries = DB::getQueryLog();

            return response()->json(['error' => 'Something went wrong', 'message' => 'Pengguna Gagal Dihapus'], 400);
        }
    }

    public function forgot(Request $request)
    {
        return view('forgot');
    }
    
    public function kirim(Request $request)
    {
        // Validate incoming request (email and password)
        $request->validate([
            'email' => 'required|email',
        ]);

        // Send login credentials to third-party API for authentication
        $response = Http::post(env('API_URL').'/forgot-password', [
            'email' => $request->email,
            'is_cms'=> 1,
        ]);
        $data = $response->object();
        
        // Check if the third-party API returned a successful response
        if ($data->message != "Success"){
        return redirect()->route('forgot')
            ->with(['error' => $data->message])
            ->withInput();
        }else{
            return redirect()->route('forgot')
            ->with(['success' => "Periksa Email anda untuk melakukan ganti password"]);
        }
    }

    public function reset(Request $request){
        $code = $request->get('code');

        $response = Http::post(env('API_URL').'/verify-token', [
            'code' => $code,
        ]);
        $data = $response->object();
        
        // Check if the third-party API returned a successful response
        if ($data->message == "Success"){
            return view('reset');
        }else{
            return view('expired');
        }
    }

    public function resetKirim(Request $request){
        $request->validate([
            'code' => 'required',
            'password' => 'required',
            'password_re' => 'required',
        ]);

        $response = Http::post(env('API_URL').'/reset-password', [
            'code' => $request->post('code'),
            'newPassword' => $request->post('password'),
            'newPasswordConfirm' => $request->post('password_re'),
        ]);
        $data = $response->object();
        
        // Check if the third-party API returned a successful response
        if ($data->message != "Success"){
            return redirect()->route('login')
                ->with(['error' => $data->message])
                ->withInput();
            }else{
                return redirect()->route('login')
                ->with(['success' => "Reset Password Berhasil, Silahkan Login Menggunakan Password yang baru"]);
            }
    }

}
