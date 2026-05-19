<?php

namespace App\Http\Controllers;

use CURLFile;
use Carbon\Carbon;
use CURLStringFile;
use Illuminate\Http\Request;
use Illuminate\Support\Facades\Http;
use Illuminate\Support\Facades\Log;
// use GuzzleHttp\Client;

class UploadController extends Controller
{
    public function upload(Request $request)
    {
    //     $input['upload_type'] = $request->query('upload-type');
    //     // $input['fileUpload'] = true;
    //     $file = $request->files->get('file');
    //     $body = fopen($file->getRealPath(), 'r');
    //     // $response = Http::withHeaders([
    //     //     'Content-Type' => 'multipart/form-data'
    //     // ])->attach(
    //     //     'file', $body, $file->getClientOriginalName(), ['Content-Type' => $file->getMimeType()]
    //     // )->post('https://project.bi.sentech.id/api/v1/stg/price', $input);
        
    //     $response = Http::attach(
    //         'file', $body, $file->getClientOriginalName(), ['Content-Type' => $file->getMimeType()]
    //         )->asMultipart()
    //     ->post(env('API_URL').'/price', [
    //         'upload-type' => $request->query("upload-type")
    //     ]);
    //     dd($response);
    //     if($response->status() == 201) {
    //         return response()->json(['success'=> $file->getClientOriginalName()]);
    //     }else{
    //         return response()->json(['error'=> $file->getClientOriginalName()]);
    //     }

        $file = $request->file('file');
        try {
            // Read file contents and name
            $fileContents = file_get_contents($file->getPathname());
            $fileName = $file->getClientOriginalName();

            // Send the file to the third-party API
            $response = Http::attach(
                'file',            // The field name expected by the API
                $fileContents,     // File contents
                $fileName          // File name
            )->asMultipart()->post(env('API_URL').'/price', [
                'upload-type' => $request->query("upload-type"), // Add any additional fields required by the API
            ]);

            // Handle the API response
            if ($response->successful()) {
                return response()->json([
                    'message' => 'File uploaded successfully!',
                    'data' => $response->json(),
                ]);
            } else {
                return response()->json([
                    'message' => 'File upload failed.',
                    'status' => $response->status(),
                    'error' => $response->body(),
                ], $response->status());
            }
        } catch (\Exception $e) {
            // Handle errors
            return response()->json([
                'message' => 'An error occurred during file upload.',
                'error' => $e->getMessage(),
            ], 500);
        }
        
    }
    
    public function neraca(Request $request)
    {
        $file = $request->file('file');
        try {
            // Read file contents and name
            $fileContents = file_get_contents($file->getPathname());
            $fileName = $file->getClientOriginalName();

            // Send the file to the third-party API
            $response = Http::attach(
                'file',            // The field name expected by the API
                $fileContents,     // File contents
                $fileName          // File name
            )->asMultipart()->post(env('API_URL').'/neraca', [
                'upload-type' => $request->query("upload-type"), // Add any additional fields required by the API
            ]);

            // Handle the API response
            if ($response->successful()) {
                return response()->json([
                    'message' => 'File uploaded successfully!',
                    'data' => $response->json(),
                ]);
            } else {
                return response()->json([
                    'message' => 'File upload failed.',
                    'status' => $response->status(),
                    'error' => $response->body(),
                ], $response->status());
            }
        } catch (\Exception $e) {
            // Handle errors
            return response()->json([
                'message' => 'An error occurred during file upload.',
                'error' => $e->getMessage(),
            ], 500);
        }
        
    }

    public function uploadHistory(Request $request)
    {
        $module = $request->get('module');
        $start = $request->get('start');
        $length = $request->get('length');
        $search = $request->input('search.value');
        if ($start == 0) {
            $page = 1;
        }else{
            $page = ($start/$length)+1;
        }
        $url = env('API_URL').'/price/upload-history?module='.$module.'&page='.$page.'&limit='.$length.'&search='.$search;
        // echo $url;
        try {
            $result = Http::get($url);
            $data = array();
            foreach($result->object()->data as $dt){
                $date = Carbon::parse($dt->createdAt);
                $cek = array(
                    "createdAt" => $date->format("Y-m-d H:i:s"),
                    "fileName" => $dt->fileName,
                    "rowTotal" => $dt->rowTotal
                );
                array_push($data, $cek);
            }
            $result  = array(
                "recordsTotal" => $result->object()->totalData,
                "recordsFiltered" => $result->object()->totalData,
                "data" => $data
            );
        }catch (Exception $e) {
            $result  = array(
                "recordsTotal" => 0,
                "recordsFiltered" => 0,
                "data" => null
            );
        }

        return response()->json($result);
    }

}
