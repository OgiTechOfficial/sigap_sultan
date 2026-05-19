<?php

namespace App\Http\Controllers;

use CURLFile;
use CURLStringFile;
use Illuminate\Http\Request;
use Illuminate\Support\Facades\Http;
// use GuzzleHttp\Client;

class JabatanController extends Controller
{
    public function index(Request $request)
    {
        return view('jabatan',[
            'active_menu' => 'jabatan'
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
        $result = Http::get(env('API_URL').'/position');
        $data = array();
        foreach($result->object()->data as $dt){
            $privileges = array();
            foreach($dt->privileges as $pr){
                array_push($privileges, $pr->menu);
            }
            $url = url('jabatan/detail/'.$dt->id);
            $cek = array(
                "name" => $dt->name,
                "privileges" => implode(", ", $privileges),
                "action" => "<a href='".$url."' class='btn btn-outline-secondary'>Ubah</a>
                <button class='btn btn-outline-secondary del-pos' data-id='".$dt->id."'>Hapus</button>"
            );
            array_push($data, $cek);
        }
        $result  = array(
            "draw" => 2,
            "recordsTotal" => 57,
            "recordsFiltered" => 57,
            "data" => $data
        );
        return $result;
    }

    
    
    public function add(Request $request)
    {
        $result = Http::get(env('API_URL').'/menu');
        return view('jabatan-add',[
            'active_menu' => 'jabatan',
            'menudata' => $result->object()->data
        ]);
    }

    public function simpan(Request $request)
    {
        // dd($request->menu);
        $kirim['position'] = $request->position;
        $kirim['privileges'] = array();
        $i=0;
        foreach($request->menu as $menu){
            $kirim['privileges'][$i]['menuId'] = (int)$menu;
            $kirim['privileges'][$i]['permissions']['read'] = 1;
            $kirim['privileges'][$i]['permissions']['create'] = 1;
            $kirim['privileges'][$i]['permissions']['update'] = 1;
            $kirim['privileges'][$i]['permissions']['delete'] = 1;
            $i++;
        }
        
        $kirimdata = [
            'position' => $request->position,
            'privileges' => $kirim['privileges'],
        ];

        $response = Http::post(env('API_URL').'/position', $kirimdata);
        
        $data = $response->object();
        if ($data->message != "Success"){
            return response()->json([
                'error' => 'Something went wrong',
                'message' => $data->message
            ], 400);
        }else{
            return response()->json(['message' => $data->message]);
        }
    }
    
    public function delete(Request $request)
    {
        $response = Http::delete(env('API_URL').'/position/'.$request->id);
        
        $data = $response->object();
        if ($data->message != "Success"){
            return response()->json([
                'error' => 'Something went wrong',
                'message' => $data->message
            ], 400);
        }else{
            return response()->json(['message' => $data->message]);
        }
    }
}
