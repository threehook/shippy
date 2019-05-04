import { Component, OnInit, Output, EventEmitter } from '@angular/core';
import { MatDialog } from '@angular/material';
// import { AddConditionsComponent, AddConditionsViewModel } from '../add-conditions/add-conditions.component';
// import { DnDService } from '../services/dnd.service';

@Component({
  selector: 'agreement-toolbox',
  templateUrl: './toolbox.component.html',
  styleUrls: ['./toolbox.component.css']
})
export class ToolboxComponent implements OnInit {

  @Output()
  // addSettlement = new EventEmitter<AddConditionsViewModel>();

  // constructor(private dialog: MatDialog, private dnd: DnDService) { }

  ngOnInit() {
  }

  addNewSettlement() {
    // const conditionsDialog = this.dialog.open(AddConditionsComponent, {
    //   width: '400px'
    // });

    // conditionsDialog.afterClosed().subscribe((result: AddConditionsViewModel) => {
    //   if (result)
    //   {
    //     this.addSettlement.emit(result);
    //   }
    // })
  }

  // dragStart(ev: DragEvent) {
  //   this.dnd.startDrag(ev);
  //   ev.dataTransfer.setData("text/plain", '');//Needed to start DnD.
  // }

}
