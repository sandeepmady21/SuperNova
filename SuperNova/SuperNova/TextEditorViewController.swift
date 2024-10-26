//
//  ContentView.swift
//  SuperNova
//
//  Created by Michael Torkaman on 6/28/24.
//

import UIKit

class TextEditorViewController: UIViewController {

    @IBOutlet weak var textView: UITextView!

    override func viewDidLoad() {
        super.viewDidLoad()
        // Additional setup after loading the view.
    }

    // Example function to save text to a file
    @IBAction func saveText(_ sender: UIBarButtonItem) {
        guard let text = textView.text else { return }
        let fileURL = getDocumentsDirectory().appendingPathComponent("savedText.txt")
        do {
            try text.write(to: fileURL, atomically: true, encoding: .utf8)
            showAlert(title: "Success", message: "Text saved successfully.")
        } catch {
            showAlert(title: "Error", message: "Failed to save text: \(error.localizedDescription)")
        }
    }

    // Helper function to get the documents directory
    func getDocumentsDirectory() -> URL {
        return FileManager.default.urls(for: .documentDirectory, in: .userDomainMask).first!
    }

    // Helper function to display alerts
    func showAlert(title: String, message: String) {
        let alertController = UIAlertController(title: title, message: message, preferredStyle: .alert)
        alertController.addAction(UIAlertAction(title: "OK", style: .default, handler: nil))
        present(alertController, animated: true, completion: nil)
    }
}
