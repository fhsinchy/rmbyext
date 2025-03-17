import unittest
import sys
from pathlib import Path
from unittest.mock import patch, mock_open
from io import StringIO
from rmbyext import main


class TestMainFunction(unittest.TestCase):
    @patch("sys.exit")
    @patch("sys.stdout", new_callable=StringIO)
    def test_no_arguments(self, mock_stdout, mock_exit):
        sys.argv = ["rmbyext.py"]
        main()
        mock_exit.assert_called_with("NO ARGUMENT PASSED.")

    @patch("sys.stdout", new_callable=StringIO)
    @patch("pathlib.Path.rglob")
    @patch("pathlib.Path.is_file", return_value=True)
    @patch("pathlib.Path.unlink")
    @patch("builtins.open", new_callable=mock_open)
    def test_file_deletion(
        self, mock_open, mock_unlink, mock_is_file, mock_rglob, mock_stdout
    ):
        sys.argv = ["rmbyext.py", "txt"]
        mock_rglob.return_value = [Path("test1.txt"), Path("test2.txt")]
        main()
        self.assertIn("Removing: TXT", mock_stdout.getvalue())
        self.assertIn("test1.txt", mock_stdout.getvalue())
        self.assertIn("test2.txt", mock_stdout.getvalue())
        mock_unlink.assert_called()
        mock_open.assert_called()

    @patch("sys.stdout", new_callable=StringIO)
    @patch("pathlib.Path.rglob", return_value=[])
    def test_no_files_to_delete(self, mock_rglob, mock_stdout):
        sys.argv = ["rmbyext.py", "txt"]
        main()
        self.assertIn("NO TXT FILES TO REMOVE.", mock_stdout.getvalue())
